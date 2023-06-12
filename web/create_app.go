package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// initialize web server with correspondant configuration and all application configurations
func CreateApp() *fiber.App {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		Views:         html.New("./templates", ".html"),
	})

	RegisterEndpoints(app)

	return app
}
