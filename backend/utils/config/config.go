package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	RedisAddress string
)

func LoadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("読み込みできませんでした: %v", err)
	}
	RedisAddress = os.Getenv("REDIS_ADDRESS")
}
