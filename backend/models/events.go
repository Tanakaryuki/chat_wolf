package models

type EventType int
type Win int

const (
	CreateRoom EventType = iota + 1
	EnterRoom
	ChangeRoomOwner
	ExitRoom
	SendChat
	SetOption
	StartGame
	SendTime
	AskQuestion
	EndQandA
	GiveAnswer
	VoteEvent
	GameResult
	PrepareCompletion
)

const (
	Wolf Win = iota + 1
	citizen
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
	VoteEnded   string `json:"vote_ended,omitempty"`
}

type Option struct {
	TurnNum         uint `json:"turn_num"`
	DiscussTime     uint `json:"discuss_time"`
	VoteTime        uint `json:"vote_time"`
	ParticipantsNum uint `json:"participantsnum"`
}
