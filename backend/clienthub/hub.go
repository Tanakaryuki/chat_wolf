package clienthub

import (
	"encoding/json"
	"math/rand"
	"strconv"

	"github.com/Tanakaryuki/chat_wolf/models"
	"github.com/Tanakaryuki/chat_wolf/redis"
	"github.com/labstack/echo/v4"
)

var RoomID int = 10000

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[string]map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan *models.Broadcast

	// Register requests from the clients.
	createRoom chan *ClientPrtocol
	enterRoom  chan *ClientPrtocol
	sendChat   chan *ClientPrtocol

	// Unregister requests from clients.
	unregister chan *Client

	logger echo.Logger
}

func InitSetDataFromProtocol(ClientPrtocol *ClientPrtocol) models.SetData {
	SetData := models.SetData{
		User:    []models.UserForRedis{},
		ChatLog: []models.ChatLog{},
		Room:    models.RoomForRedis{},
		Option:  models.Option{},
	}
	SetData.User = append(SetData.User, models.UserForRedis{
		ID:          ClientPrtocol.Protocol.ID,
		Conn:        ClientPrtocol.Client.Conn,
		DisplayName: ClientPrtocol.Protocol.DisplayName,
		Icon:        ClientPrtocol.Protocol.Icon,
		IsWolf:      ClientPrtocol.Protocol.IsWolf,
		Score:       ClientPrtocol.Protocol.Score,
		Word:        ClientPrtocol.Protocol.Word,
		Vote: models.Vote{
			ID:          ClientPrtocol.Protocol.ID,
			DisplayName: ClientPrtocol.Protocol.DisplayName,
		},
	})
	SetData.ChatLog = append(SetData.ChatLog, models.ChatLog{
		User: models.UserForChatLog{
			ID:          ClientPrtocol.Protocol.ID,
			DisplayName: ClientPrtocol.Protocol.DisplayName,
			Icon:        ClientPrtocol.Protocol.Icon,
		},
		ChatText: ClientPrtocol.Protocol.ChatText,
	})
	SetData.Room = models.RoomForRedis{
		RoomOwnerID: ClientPrtocol.Protocol.RoomOwnerID,
		VoteEnded:   ClientPrtocol.Protocol.VoteEnded,
	}
	SetData.Option = models.Option{
		TurnNum:         ClientPrtocol.Protocol.TurnNum,
		DiscussTime:     ClientPrtocol.Protocol.DiscussTime,
		VoteTime:        ClientPrtocol.Protocol.VoteTime,
		ParticipantsNum: ClientPrtocol.Protocol.ParticipantsNum,
	}
	return SetData
}

func NewHub(logger echo.Logger) *Hub {
	return &Hub{
		broadcast:  make(chan *models.Broadcast),
		enterRoom:  make(chan *ClientPrtocol),
		createRoom: make(chan *ClientPrtocol),
		unregister: make(chan *Client),
		clients:    make(map[string]map[*Client]bool),
		logger:     logger,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.createRoom:
			RoomID += rand.Intn(20)
			roomnum := strconv.Itoa(RoomID)
			if _, ok := h.clients[roomnum][&client.Client]; ok {
				h.clients[roomnum][&client.Client] = true
			} else {
				h.clients[roomnum] = map[*Client]bool{}
				h.clients[roomnum][&client.Client] = true
			}
			client.Protocol.Room.RoomID = roomnum
			setData := InitSetDataFromProtocol(client)
			data, err := json.Marshal(setData)
			if err != nil {
				h.logger.Error(err)
				break
			}
			redis.Set(roomnum, data)
			data, err = json.Marshal(client.Protocol)
			if err != nil {
				h.logger.Error(err)
				break
			}
			h.logger.Error("start broadcast")
			h.Broadcast(roomnum, data)
		case client := <-h.enterRoom:
			getData, err := redis.Get(client.Protocol.Room.RoomID)
			if err != nil {
				h.logger.Error(err)
				break
			}
			var setData models.SetData
			err = json.Unmarshal(getData, &setData)
			if err != nil {
				h.logger.Error(err)
				break
			}
			h.logger.Error(setData.User)
			setData.User = append(setData.User, models.UserForRedis{
				ID:          client.Protocol.User.ID,
				Conn:        client.Client.Conn,
				DisplayName: client.Protocol.User.DisplayName,
				Icon:        client.Protocol.User.Icon,
			})
			data, err := json.Marshal(setData)
			if err != nil {
				h.logger.Error(err)
				break
			}
			if err = redis.Set(client.Protocol.Room.RoomID, data); err != nil {
				h.logger.Error(err)
				break
			}
			roomNum := client.Protocol.Room.RoomID
			if _, ok := h.clients[roomNum][&client.Client]; ok {
				h.clients[roomNum][&client.Client] = true
			} else {
				h.clients[roomNum] = map[*Client]bool{}
				h.clients[roomNum][&client.Client] = true
			}
			data, err = json.Marshal(client.Protocol)
			if err != nil {
				h.logger.Error(err)
				break
			}
			h.Broadcast(client.Protocol.Room.RoomID, data)
		case client := <-h.sendChat:
			getData, err := redis.Get(client.Protocol.Room.RoomID)
			if err != nil {
				h.logger.Error(err)
				break
			}
			var setData *models.SetData
			err = json.Unmarshal(getData, setData)
			if err != nil {
				h.logger.Error(err)
				break
			}

			setData.ChatLog = append(setData.ChatLog, models.ChatLog{
				User: models.UserForChatLog{
					ID:          client.Protocol.User.ID,
					DisplayName: client.Protocol.User.DisplayName,
					Icon:        client.Protocol.User.Icon,
				},
				ChatText: client.Protocol.ChatText,
			})
			data, err := json.Marshal(setData)
			if err != nil {
				h.logger.Error(err)
				break
			}
			redis.Set(client.Protocol.Room.RoomID, data)
			data, err = json.Marshal(client.Protocol)
			if err != nil {
				h.logger.Error(err)
				break
			}
			h.Broadcast(client.Protocol.Room.RoomID, data)
		case broadcast := <-h.broadcast:
			for client := range h.clients[broadcast.RoomNum] {
				select {
				case client.Send <- broadcast.Data:
				default:
					close(client.Send)
					delete(h.clients[broadcast.RoomNum], client)
				}
			}
		}
	}
}

func (h *Hub) Broadcast(roomNum string, data []byte) {
	h.logger.Error("start broadcasts 2")
	h.logger.Error(h.clients[roomNum])
	for client := range h.clients[roomNum] {
		select {
		case client.Send <- data:
		default:
			close(client.Send)
			delete(h.clients[roomNum], client)
		}
	}
}
