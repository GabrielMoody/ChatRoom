package ws

import (
	"fmt"

	"github.com/GabrielMoody/chat-app/server/internal/dto"
	"github.com/GabrielMoody/chat-app/server/internal/mysql"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Message struct {
	Username string
	RoomID   string
	Message  string
}

type Room struct {
	RoomID string
	Name   string
	Client map[string]*Client
}

type Handler struct {
	db  *gorm.DB
	hub *Hub
}

func (h *Handler) CreateRoom(c *fiber.Ctx) error {
	var r dto.RoomReq

	jwtUser := c.Locals("user").(*jwt.Token)
	claims := jwtUser.Claims.(jwt.MapClaims)
	username := claims["user"].(string)
	id := claims["ID"].(string)

	if err := c.BodyParser(&r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	room := mysql.Room{
		ID:        uuid.NewString(),
		Name:      r.Name,
		CreatedBy: username,
	}

	user := mysql.User{
		ID:     id,
		RoomID: []mysql.Room{room},
	}

	if err := h.db.Create(&room).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.db.Updates(user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	h.hub.Rooms[room.ID] = &Room{
		RoomID: room.ID,
		Name:   room.Name,
		Client: make(map[string]*Client),
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"rooms": room,
	})
}

func (h *Handler) GetJoinedRoom(c *fiber.Ctx) error {
	id := c.Params("userid")

	var user mysql.User

	if err := h.db.WithContext(c.Context()).Preload("rooms").First(&user, "id = ? ", id).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"rooms": user.RoomID,
	})
}

func (h *Handler) FindRoom(c *fiber.Ctx) error {
	var data []mysql.Room

	room := c.Query("room")

	if err := h.db.Find(&data, "name LIKE ?", fmt.Sprintf("%%%s%%", room)).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"rooms": data,
	})
}

func (h *Handler) JoinRoom(ctx *fiber.Ctx) func(*websocket.Conn) {
	return func(c *websocket.Conn) {
		roomID := c.Params("roomID")
		username := c.Cookies("X-Username")

		client := &Client{
			Conn:     c,
			Username: username,
			RoomID:   roomID,
		}

		h.hub.Register <- client

	}
}

func SendMessage(h *Hub) func(*websocket.Conn) {
	return func(c *websocket.Conn) {
		defer func() {
			h.Removal <- c
			_ = c.Close()
		}()

		//from := c.Cookies("X-Username")

		//h.Register <- &Room{
		//	RoomID: uuid.NewString(),
		//	Name: from,
		//	Client:
		//}

		for {
			msgType, msg, err := c.ReadMessage()
			if err != nil {
				return
			}

			if msgType == websocket.TextMessage {
				h.Message <- &Message{
					Message: string(msg),
				}
			}
		}

	}
}

func NewHandler(db *gorm.DB, hub *Hub) *Handler {
	return &Handler{
		db:  db,
		hub: hub,
	}
}
