package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Tanakaryuki/chat_wolf/clienthub"
	"github.com/Tanakaryuki/chat_wolf/redis"
	"github.com/Tanakaryuki/chat_wolf/utils/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var ctx = context.Background()

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	config.LoadEnv()

	redis.SetupRedis()
	hub := clienthub.NewHub(e.Logger)
	go hub.Run()
	ping, err := redis.Cache.Ping(ctx).Result()
	if err != nil {
		log.Printf("Could not connect to Redis: %v", err)
	}
	fmt.Println(ping)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/ws", func(c echo.Context) error {
		clienthub.ServeWs(hub, c)
		return nil
	})
	e.GET("/hello", clienthub.Hello) //WebSocketテスト用
	e.Logger.Fatal(e.Start(":8080"))
}
