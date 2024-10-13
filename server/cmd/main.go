package main

import (
	"github.com/GabrielMoody/chat-app/server/internal"
	"github.com/GabrielMoody/chat-app/server/internal/mysql"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8000, http://localhost:63342",
		AllowCredentials: true,
	}))

	db := mysql.NewConnection()
	api := app.Group("/api/v1")

	internal.ChatRoutes(api, db)

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}

		return fiber.ErrUpgradeRequired
	})

	err := app.Listen("localhost:8000")

	if err != nil {
		panic(err.Error())
	}

}
