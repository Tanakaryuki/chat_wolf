package models

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
