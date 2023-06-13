package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// initialize web server with correspondant configuration and all application configurations
func CreateApp() *fiber.App {

	// defining html rendering engine and configuring its behaviour
	engine := html.New("./templates", ".html")

	// create fiber app and append endpoints
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		Views:         engine,
	})

	RegisterEndpoints(app)

	return app
}
