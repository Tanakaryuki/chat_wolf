package json

import (
	"encoding/json"
	"log"

	"github.com/Tanakaryuki/chat_wolf/models"
)

func JsonToString(jsonMessage models.Protocol) string {
	bytes, err := json.Marshal(jsonMessage)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}
