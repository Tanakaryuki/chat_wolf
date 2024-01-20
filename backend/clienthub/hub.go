package clienthub

import (
	"encoding/json"
	"math/rand"
	"strconv"

	"github.com/Tanakaryuki/chat_wolf/models"
	"github.com/Tanakaryuki/chat_wolf/redis"
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
}

func SetDataFromProtocol(ClientPrtocol *ClientPrtocol) models.SetData {
	var SetData models.SetData
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

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan *models.Broadcast),
		enterRoom:  make(chan *ClientPrtocol),
		createRoom: make(chan *ClientPrtocol),
		unregister: make(chan *Client),
		clients:    make(map[string]map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.createRoom:
			RoomID += rand.Intn(20)
			roomnum := strconv.Itoa(RoomID)
			h.clients[roomnum][&client.Client] = true
			setData := SetDataFromProtocol(client)
			data, _ := json.Marshal(setData)
			redis.Set(roomnum, data)
		case client := <-h.enterRoom:
			getData, err := redis.Get(client.Protocol.Room.RoomID)
			if err == nil {
				var setData *models.SetData
				json.Unmarshal(getData, setData)
				h.clients[client.Protocol.Room.RoomID][&client.Client] = true
				data, _ := json.Marshal(client.Protocol)
				h.broadcast <- &models.Broadcast{
					RoomNum: client.Protocol.Room.RoomID,
					Data:    data,
				}
			}
		case client := <-h.sendChat:
			getData, err := redis.Get(client.Protocol.Room.RoomID)
			if err == nil {
				var setData *models.SetData
				json.Unmarshal(getData, setData)
				setData.ChatLog = append(setData.ChatLog, models.ChatLog{
					User: models.UserForChatLog{
						ID:          client.Protocol.User.ID,
						DisplayName: client.Protocol.User.DisplayName,
						Icon:        client.Protocol.User.Icon,
					},
					ChatText: client.Protocol.ChatText,
				})
				data, _ := json.Marshal(setData)
				redis.Set(client.Protocol.Room.RoomID, data)
				data, _ = json.Marshal(client.Protocol)
				h.broadcast <- &models.Broadcast{
					RoomNum: client.Protocol.Room.RoomID,
					Data:    data,
				}
			}
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
