package middleware

import (
	"github.com/GabrielMoody/chat-app/server/internal/helper"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware() fiber.Handler {
	v := helper.LoadEnv()
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(v.GetString("JWT_SECRET"))},
	})
}
