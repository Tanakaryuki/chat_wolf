package utils

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

func StringToJson(rawMsg string) models.SetData {
	jsonBytes := []byte(rawMsg)
	var jsonStruct models.SetData
	if err := json.Unmarshal(jsonBytes, &jsonStruct); err != nil {
		log.Fatal(err)
	}
	return jsonStruct
}
