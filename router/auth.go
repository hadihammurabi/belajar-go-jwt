package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hadihammurabi/belajar-go-jwt/controller"
)

func Auth(app *fiber.App) {
	ctr := controller.NewAuthController()

	group := app.Group("/auth")
	group.Post("login", ctr.Login)
	group.Get("info", ctr.Info)
}
