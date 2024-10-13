// package internal

// import (
// 	"log"

// 	"github.com/gofiber/contrib/websocket"
// )

// type message struct {
// 	From string
// 	To   string
// 	Msg  string
// }

// type PrivateConnection struct {
// 	Clients map[*websocket.Conn]*websocket.Conn
// 	Message chan message
// }

// func NewHub() *PrivateConnection {
// 	return &PrivateConnection{
// 		Clients: make(map[*websocket.Conn]*websocket.Conn),
// 		Message: make(chan message),
// 	}
// }

// func (p *PrivateConnection) Run() {
// 	for {
// 		select {
// 		case msg := <-p.Message:
// 			p.Clients[]
// 		}
// 	}
// }

// func SendPrivateMessage(p *PrivateConnection) func(*websocket.Conn) {
// 	return func(c *websocket.Conn) {
// 		user := c.Query("user")
// 		to := c.Query("to")

// 		p.WS <- c

// 		msgType, msg, err := c.ReadMessage()

// 		for {
// 			if err != nil {
// 				log.Fatal(err.Error())
// 				return
// 			}

// 			if msgType == websocket.TextMessage {
// 				p.Msg <- message{
// 					From: user,
// 					To:   to,
// 					Msg:  string(msg),
// 				}
// 			}
// 		}
// 	}
// }
