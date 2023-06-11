package main

import (
	"github.com/carepollo/librecode/utils"
	"github.com/carepollo/librecode/web"
	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.LoadEnv()
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
	})
	web.RegisterEndpoints(app)
	app.Listen(":8080")
}
