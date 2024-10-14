package internal

import (
	"github.com/GabrielMoody/chat-app/server/internal/middleware"
	"github.com/GabrielMoody/chat-app/server/internal/ws"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ChatRoutes(r fiber.Router, db *gorm.DB) {
	hub := ws.Hub{
		Rooms:    make(map[string]*ws.Room),
		Register: make(chan *ws.Client),
		Removal:  make(chan *websocket.Conn),
		Message:  make(chan *ws.Message),
	}

	userHandler := NewUserHandler(db)
	chatHandler := ws.NewHandler(db, &hub)
	api := r.Group("/")

	go hub.Run()

	api.Post("/register", userHandler.CreateUser)
	api.Post("/login", userHandler.LoginUser)
	api.Get("/rooms", middleware.JWTMiddleware(), chatHandler.FindRoom)
	api.Post("/rooms", middleware.JWTMiddleware(), chatHandler.CreateRoom)

	api.Use("/ws/chat", middleware.JWTMiddleware(), websocket.New(ws.SendMessage(&hub)))

}
