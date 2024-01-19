package main

import (
	"net/http"

	"github.com/Tanakaryuki/chat_wolf/utils/config"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

var Cache *redis.Client

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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
