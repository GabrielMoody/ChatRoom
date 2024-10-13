package internal

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ChatRoutes(r fiber.Router, db *gorm.DB) {
	userHandler := NewUserHandler(db)
	api := r.Group("/")

	// hub := ws.NewHub()

	//go hub.Run()

	api.Post("/register", userHandler.CreateUser)
	api.Post("/login", userHandler.LoginUser)

	// api.Use("/ws/chat", websocket.New(ws.SendMessage(hub)))
}
