package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Tanakaryuki/chat_wolf/utils/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())

	config.LoadEnv()

	setupRedis()

	hub := NewHub()
	go hub.run()
	ping, err := Cache.Ping(ctx).Result()
	if err != nil {
		log.Printf("Could not connect to Redis: %v", err)
	}
	fmt.Println(ping)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/ws", func(c echo.Context) error {
		ServeWs(hub, c.Response().Writer, c.Request())
		return nil
	})
	e.GET("/hello", Hello) //WebSocketテスト用
	e.Logger.Fatal(e.Start(":8080"))
}
