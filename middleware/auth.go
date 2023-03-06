package middleware

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthBearer(tokenSigningKey []byte) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		authorization := ctx.Get("Authorization")
		if authorization == "" {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		bearer := strings.Split(authorization, " ")
		if len(bearer) < 2 {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		token, err := jwt.ParseWithClaims(bearer[1], jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
			return tokenSigningKey, nil
		})
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		ctx.Locals("token", token)

		return ctx.Next()
	}
}
