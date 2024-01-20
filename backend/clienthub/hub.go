package clienthub

import (
	"math/rand"
	"strconv"

	"github.com/Tanakaryuki/chat_wolf/models"
)

var RoomID int = 10000

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[string]map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	createRoom chan *ClientPrtocol
	enterRoom  chan *Client

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
		broadcast:  make(chan []byte),
		enterRoom:  make(chan *Client),
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
			SetDataFromProtocol(client)
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
