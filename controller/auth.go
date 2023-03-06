package controller

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hadihammurabi/belajar-go-jwt/model"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

var TokenSigningKey = []byte("surpriselySecretDuaarrrrr...")

func (controller AuthController) Login(ctx *fiber.Ctx) error {
	tokenData := jwt.MapClaims{
		model.TokenClaimUserID: 10,
		model.TokenClaimExp:    time.Now().Add(30 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenData)
	tokenStr, err := token.SignedString(TokenSigningKey)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"token": tokenStr,
	})
}

func (controller AuthController) Info(ctx *fiber.Ctx) error {
	token := ctx.Locals("token").(*jwt.Token)
	tokenData, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

	return ctx.JSON(fiber.Map{
		model.TokenClaimUserID: tokenData[model.TokenClaimUserID],
	})
}
