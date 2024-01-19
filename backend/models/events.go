package models

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
	EndQandA          = EventType("end_q_and_a")
	GiveAnswer        = EventType("give_answer")
	VoteEvent         = EventType("vote_event")
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
	VoteEnded   string `json:"vote_ended,omitempty"`
}

type Option struct {
	TurnNum         uint `json:"turn_num"`
	DiscussTime     uint `json:"discuss_time"`
	VoteTime        uint `json:"vote_time"`
	ParticipantsNum uint `json:"participantsnum"`
}

type Users struct {
	DisplayName string `json:"display_name,omitempty"`
	IsWolf      bool   `json:"is_wolf"`
	Score       uint   `json:"score"`
	Word        string `json:"word,omitempty"`
	Vote        `json:"vote,omitempty"`
}

type Protocol struct {
	EventType EventType `json:"event_type"`
	User      `json:"user,omitempty"`
	Room      `json:"room,omitempty"`
	ChatText  string `json:"chat_text,omitempty"`
	Option    `json:"option,omitempty"`
	TimeNow   int     `json:"time_now,omitempty"`
	Win       Win     `json:"win,omitempty"`
	Users     []Users `json:"users,omitempty"`
}
