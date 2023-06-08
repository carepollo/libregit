package main

import (
	"github.com/carepollo/librecode/web"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	web.RegisterEndpoints(app)
	app.Listen(":8080")
}
