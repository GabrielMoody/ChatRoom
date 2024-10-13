package ws

import "github.com/gofiber/contrib/websocket"

type Hub struct {
	Rooms    map[string]*Room
	Register chan *Client
	Removal  chan *websocket.Conn
	Message  chan *Message
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				r := h.Rooms[cl.RoomID]

				if _, ok := r.Client[cl.Username]; !ok {
					r.Client[cl.Username] = cl
				}
			}

		// case cl := <-h.Removal:
		// 	delete(h.)

		case msg := <-h.Message:
			for _, user := range h.Rooms {
				if msg.RoomID == user.RoomID {
					for _, conn := range user.Client {
						conn.Conn.WriteJSON(msg)
					}
				}
			}
		}
	}
}
