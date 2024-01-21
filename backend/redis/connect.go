package redis

import (
	"github.com/Tanakaryuki/chat_wolf/utils/config"
	"github.com/redis/go-redis/v9"
)

var Cache *redis.Client

func SetupRedis() {
	Cache = redis.NewClient(&redis.Options{
		// docker-compose.ymlに指定したservice名+port
		Addr:     config.RedisAddress,
		Password: "",
		DB:       0,
	})
}
