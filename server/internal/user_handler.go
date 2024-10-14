package internal

import (
	"time"

	"github.com/GabrielMoody/chat-app/server/internal/dto"
	helper2 "github.com/GabrielMoody/chat-app/server/internal/helper"
	"github.com/GabrielMoody/chat-app/server/internal/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func (u *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user dto.UserRegistration

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := helper2.Validate.Struct(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	data := mysql.User{
		ID:       uuid.NewString(),
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashed),
	}

	if err := u.db.Create(&data).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": data,
	})
}

func (u *UserHandler) LoginUser(c *fiber.Ctx) error {
	var user dto.UserLogin
	var data mysql.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := helper2.Validate.Struct(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := u.db.Where("email = ?", user.Email).First(&data).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(user.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	claims := jwt.MapClaims{
		"ID":   data.ID,
		"user": data.Name,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	v := helper2.LoadEnv()
	t, errToken := token.SignedString([]byte(v.GetString("JWT_SECRET")))

	if errToken != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errToken.Error(),
		})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "X-Username"
	cookie.Value = data.Name
	cookie.Domain = "localhost"
	cookie.HTTPOnly = true
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(time.Hour * 24)

	c.Cookie(cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": t,
	})
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		db: db,
	}
}
