package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Tanakaryuki/chat_wolf/utils/config"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

var Cache *redis.Client
var ctx = context.Background()

func setupRedis() {
	Cache = redis.NewClient(&redis.Options{
		// docker-compose.ymlに指定したservice名+port
		Addr: config.RedisAddress,
		DB:   0,
	})
}

func main() {
	e := echo.New()

	config.LoadEnv()

	setupRedis()

	ping, err := Cache.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println(ping)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
