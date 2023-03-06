package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hadihammurabi/belajar-go-jwt/router"
)

func main() {

	app := fiber.New()
	router.Auth(app)
	app.Listen(":8080")

}
