package main

import (
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

var Cache *redis.Client

func setupRedis() {
	Cache = redis.NewClient(&redis.Options{
		// docker-compose.ymlに指定したservice名+port
		Addr: "redis:6379",
		DB:   0,
	})
}

func main() {
	e := echo.New()

	setupRedis()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}