package models

import "github.com/gorilla/websocket"

type EventType string
type Win string

const (
	CreateRoom        = EventType("create_room")
	EnterRoom         = EventType("enter_room")
	ChangeRoomOwner   = EventType("change_room_owner")
	ExitRoom          = EventType("exit_room")
	SendChat          = EventType("send_chat")
	SetOption         = EventType("set_option")
	StartGame         = EventType("start_game")
	SendTime          = EventType("send_time")
	AskQuestion       = EventType("ask_question")
	EndQandA          = EventType("end_Q_and_A")
	GiveAnswer        = EventType("give_answer")
	VoteEvent         = EventType("vote")
	GameResult        = EventType("game_result")
	PrepareCompletion = EventType("prepare_completion")
)

const (
	Wolf    = Win("wolf")
	citizen = Win("citizen")
)

type Vote struct {
	ID          string `json:"id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

type User struct {
	ID            string `json:"id,omitempty"`
	DisplayName   string `json:"display_name,omitempty"`
	Icon          string `json:"icon,omitempty"`
	IsWolf        bool   `json:"is_wolf"`
	Score         uint   `json:"score"`
	Word          string `json:"word,omitempty"`
	IsParticipant bool   `json:"is_participant"`
	Vote          `json:"vote,omitempty"`
}

type Room struct {
	RoomID      string `json:"room_id,omitempty"`
	RoomOwnerID string `json:"room_owner_id,omitempty"`
	VoteEnded   bool   `json:"vote_ended,omitempty"`
}

type Option struct {
	TurnNum         uint `json:"turn_num"`
	DiscussTime     uint `json:"discuss_time"`
	VoteTime        uint `json:"vote_time"`
	ParticipantsNum uint `json:"participantsnum"`
}

type Protocol struct {
	EventType EventType `json:"event_type"`
	User      `json:"user,omitempty"`
	Room      `json:"room,omitempty"`
	ChatText  string `json:"chat_text,omitempty"`
	Option    `json:"option,omitempty"`
	TimeNow   int    `json:"time_now,omitempty"`
	Win       Win    `json:"win,omitempty"`
	Users     []User `json:"users,omitempty"`
}

type UserForRedis struct {
	ID            string          `json:"id,omitempty"`
	Conn          *websocket.Conn `json:"conn"`
	DisplayName   string          `json:"display_name,omitempty"`
	Icon          string          `json:"icon,omitempty"`
	IsWolf        bool            `json:"is_wolf"`
	Score         uint            `json:"score"`
	Word          string          `json:"word,omitempty"`
	IsParticipant bool            `json:"is_participant"`
	Vote          `json:"vote,omitempty"`
}

type RoomForRedis struct {
	RoomOwnerID string `json:"room_owner_id"`
	VoteEnded   bool   `json:"vote_ended"`
}

type UserForChatLog struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
	Icon        string `json:"icon"`
}

type ChatLog struct {
	User     UserForChatLog `json:"user"`
	ChatText string         `json:"chat_text"`
}

type SetData struct {
	User    []UserForRedis `json:"user"`
	ChatLog []ChatLog      `json:"chat_log"`
	Room    RoomForRedis   `json:"room"`
	Option  Option         `json:"option"`
}

type Broadcast struct {
	RoomNum string
	Data    []byte
}
