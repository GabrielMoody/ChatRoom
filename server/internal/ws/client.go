package ws

import "github.com/gofiber/contrib/websocket"

type Client struct {
	Conn     *websocket.Conn
	Username string
	RoomID   string
}
