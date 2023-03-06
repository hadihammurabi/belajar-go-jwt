package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hadihammurabi/belajar-go-jwt/controller"
	"github.com/hadihammurabi/belajar-go-jwt/middleware"
)

func Auth(app *fiber.App) {
	ctr := controller.NewAuthController()

	group := app.Group("/auth")
	group.Post("login", ctr.Login)

	group.Use(middleware.AuthBearer(controller.TokenSigningKey))
	group.Get("info", ctr.Info)
}
