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
		Addr:     config.RedisAddress,
		Password: "",
		DB:       0,
	})
}

func main() {
	e := echo.New()

	config.LoadEnv()

	setupRedis()

	hub := NewHub()
	go hub.run()
	ping, err := Cache.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println(ping)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(hub, w, r)
	})
	e.GET("/ws", Hello) //WebSocketテスト用
	e.Logger.Fatal(e.Start(":8080"))
}
