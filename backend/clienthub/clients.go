package clienthub

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Tanakaryuki/chat_wolf/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 4096
)

var (
	newline = []byte{'\n'}
	// space  = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	Hub *Hub

	// The websocket connection.
	Conn *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte
}

type ClientPrtocol struct {
	Client   Client
	Protocol models.Protocol
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump(logger echo.Logger) {
	defer func() {
		c.Hub.unregister <- c
		c.Conn.Close()
	}()
	c.Conn.SetReadLimit(maxMessageSize)
	c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		logger.Error("hello, client: %v", c)
		var message models.Protocol
		err := c.Conn.ReadJSON(&message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Error("error: %v", err)
			}
			logger.Error(err)
			break
		}
		logger.Error("start switch statment")
		cp := &ClientPrtocol{
			Client:   *c,
			Protocol: message,
		}
		logger.Error(cp)
		switch message.EventType {
		case models.CreateRoom:
			logger.Error("Call Create Room")
			c.Hub.createRoom <- cp
		case models.EnterRoom:
			c.Hub.enterRoom <- cp
		case models.ChangeRoomOwner:
		case models.ExitRoom:
		case models.SendChat:
			c.Hub.sendChat <- cp
		case models.SetOption:
		case models.StartGame:
			c.Hub.startGame <- cp
		case models.SendTime:
		case models.AskQuestion:
		case models.EndQandA:
		case models.GiveAnswer:
		case models.VoteEvent:
		case models.GameResult:
		case models.PrepareCompletion:
		default:
			logger.Error("This is Called")
		}
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump(logger echo.Logger) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				logger.Error("c.Send not ok")
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				logger.Error(err)
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				logger.Error(err)
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				logger.Error(err)
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWs(hub *Hub, c echo.Context) {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil) //c.Response().Writerになる可能性あり
	if err != nil {
		c.Logger().Error(err)
		return
	}
	client := &Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump(c.Logger())
	go client.readPump(c.Logger())
}

func Hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil) //c.Response().Writerになる可能性あり
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		fmt.Printf("%s\n", msg)
	}
}
