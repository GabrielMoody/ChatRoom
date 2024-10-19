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
		AllowOrigins:     "http://localhost:8000, http://localhost:8080",
		AllowMethods:     "GET, POST, DELETE, PUT, PATCH",
		AllowCredentials: true,
	}))

	db := mysql.NewConnection()
	api := app.Group("/api/v1")

	internal.ChatRoutes(api, db)

	app.Options("*", func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		return c.SendStatus(fiber.StatusNoContent)
	})

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
