package controller

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hadihammurabi/belajar-go-jwt/model"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

var tokenKey = []byte("surpriselySecretDuaarrrrr...")

func (controller AuthController) Login(ctx *fiber.Ctx) error {
	tokenData := jwt.MapClaims{
		model.TokenClaimUserID: 10,
		model.TokenClaimExp:    time.Now().Add(30 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenData)
	tokenStr, err := token.SignedString(tokenKey)
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
		return tokenKey, nil
	})
	if err != nil {
		fmt.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

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
