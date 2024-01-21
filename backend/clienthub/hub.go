package clienthub

import (
	"bytes"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Tanakaryuki/chat_wolf/models"
	"github.com/Tanakaryuki/chat_wolf/redis"
	"github.com/labstack/echo/v4"
)

var RoomID int = 10000

const (
	DefaultTime = 180 * time.Second
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[string]map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan *models.Broadcast

	// Register requests from the clients.
	createRoom  chan *ClientPrtocol
	enterRoom   chan *ClientPrtocol
	sendChat    chan *ClientPrtocol
	startGame   chan *ClientPrtocol
	askQuestion chan *ClientPrtocol
	voteEvent   chan *ClientPrtocol

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
		broadcast:   make(chan *models.Broadcast),
		enterRoom:   make(chan *ClientPrtocol),
		createRoom:  make(chan *ClientPrtocol),
		sendChat:    make(chan *ClientPrtocol),
		startGame:   make(chan *ClientPrtocol),
		askQuestion: make(chan *ClientPrtocol),
		voteEvent:   make(chan *ClientPrtocol),
		unregister:  make(chan *Client),
		clients:     make(map[string]map[*Client]bool),
		logger:      logger,
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
			h.Broadcast(roomnum, data)
		case client := <-h.enterRoom:
			getData, err := redis.Get(client.Protocol.Room.RoomID)
			if err != nil {
				h.logger.Error(err)
				break
			}
			h.clients[client.Protocol.Room.RoomID][&client.Client] = true
			var setData models.SetData
			err = json.Unmarshal(getData, &setData)
			if err != nil {
				h.logger.Error(err)
				break
			}
			if len(setData.User) >= 7 {
				setData.User = append(setData.User, models.UserForRedis{
					ID:            client.Protocol.User.ID,
					Conn:          client.Client.Conn,
					DisplayName:   client.Protocol.User.DisplayName,
					Icon:          client.Protocol.User.Icon,
					IsParticipant: false,
				})
			} else {
				setData.User = append(setData.User, models.UserForRedis{
					ID:            client.Protocol.User.ID,
					Conn:          client.Client.Conn,
					DisplayName:   client.Protocol.User.DisplayName,
					Icon:          client.Protocol.User.Icon,
					IsParticipant: true,
				})
			}
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
			h.logger.Error("call sendchat")
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
			h.logger.Error("call broadcast")
			h.Broadcast(client.Protocol.Room.RoomID, data)
		case client := <-h.startGame:
			roomID := client.Protocol.Room.RoomID
			ticker := time.NewTicker(1 * time.Second)
			timeout := time.After(DefaultTime)
			go func(roomID string) {
				countDown := int(DefaultTime / time.Second)
				for {
					select {
					case <-ticker.C:
						countDown--
						client.Protocol.TimeNow = countDown
						data, err := json.Marshal(client.Protocol)
						if err != nil {
							h.logger.Error(err)
							continue
						}
						h.Broadcast(roomID, data)
					case <-timeout:
						client.Protocol.TimeNow = 0
						data, err := json.Marshal(client.Protocol)
						if err != nil {
							h.logger.Error(err)
						}
						h.Broadcast(roomID, data)
						ticker.Stop()
						return
					}
				}

			}(roomID)
		case client := <-h.askQuestion:
			_, err := redis.Get(client.Protocol.Room.RoomID)
			if err != nil {
				h.logger.Error(err)
				break
			}
			data, err := json.Marshal(client.Protocol)
			if err != nil {
				h.logger.Error(err)
				break
			}
			h.Broadcast(client.Protocol.Room.RoomID, data)
			go func(client *ClientPrtocol) {
				//GPTに質問送信→質問が帰ってくる
				//質問をclient.ProtocolのchatTextに上書き
				//eventtypeの変更
				//もう一度client.ProtocolをMarshalしてdataに保存
				//broadCastを同様の引数で再呼び出し
				apiURL := "https://api.openai.com/v1/chat/completions"

				// 送信するデータ（JSONデータを仮定）
				requestBody := models.QuestionRequest{
					Model: "gpt-3.5-turbo",
					Messages: []struct {
						Role    string `json:"role"`
						Content string `json:"content"`
					}{
						{"system", "userが「りんご」について質問するので「はい」か「いいえ」か「分からない」のどれかで回答してください。「はい」か「いいえ」で答えられない質問や質問になっていないものなどは全て「分からない」で回答してください。回答は一般的な常識で行ってください。「はい」「いいえ」「分からない」以外で回答をすると無実の人間が被害に遭うので「はい」「いいえ」「分からない」以外は絶対に発言しないでください"},
						{"user", "それは赤いですか？"},
					},
				}

				jsonData, err := json.Marshal(requestBody)
				if err != nil {
					h.logger.Error(err)
					return
				}

				// HTTPリクエストの作成
				req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
				if err != nil {
					h.logger.Error(err)
					return
				}

				// ヘッダーの追加
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("Authorization", os.Getenv("SECRET_KEY"))

				// HTTPクライアントの生成
				clients := &http.Client{}

				// HTTPリクエストの送信
				response, err := clients.Do(req)
				if err != nil {
					h.logger.Error(err)
					return
				}
				defer response.Body.Close()

				// レスポンスのボディを読み取り
				body, err := io.ReadAll(response.Body)
				if err != nil {
					h.logger.Error(err)
					return
				}

				//ChatCompletion構造体にUnmarshal
				var chatCompletion models.QuestionResponse
				err = json.Unmarshal(body, &chatCompletion)
				if err != nil {
					h.logger.Error(err)
					return
				}

				// "content" フィールドを表示
				h.logger.Error(string(body))

				client.Protocol.ChatText = chatCompletion.Choices[0].Message.Content
				client.Protocol.EventType = models.GiveAnswer
				data, err := json.Marshal(client.Protocol)
				if err != nil {
					h.logger.Error(err)
				}
				h.Broadcast(client.Protocol.Room.RoomID, data)
			}(client)
		case client := <-h.voteEvent:
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
			client.Protocol.Room.VoteEnded = true
			for i := 0; i < len(setData.User); i++ {
				if client.Protocol.User.Vote.ID == setData.User[i].ID {
					setData.User[i].Vote = models.Vote{
						ID:          client.Protocol.User.Vote.ID,
						DisplayName: client.Protocol.User.Vote.DisplayName,
					}

				}
				if setData.User[i].Vote.ID == "" {
					client.Protocol.Room.VoteEnded = false
				}
			}
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
