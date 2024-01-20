package json

import (
	"encoding/json"
	"log"

	"github.com/Tanakaryuki/chat_wolf/models"
)

func StringToJson(rawMsg string) models.SetData {
	jsonBytes := []byte(rawMsg)
	var jsonMsg models.SetData
	if err := json.Unmarshal(jsonBytes, &jsonMsg); err != nil {
		log.Fatal(err)
	}
	return jsonMsg
}
