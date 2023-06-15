package web

import (
	"os"

	"github.com/carepollo/librecode/git"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// initialize web server with correspondant configuration and all application configurations
// assign to package variables some environment variables
func CreateApp() *fiber.App {
	git.GitPath = os.Getenv("GIT_ROOT")

	// defining html rendering engine and configuring its behaviour
	engine := html.New("./templates", ".html")
	engine.Reload(os.Getenv("PRODUCTION") == "true")

	// create fiber app and append endpoints
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		Views:         engine,
	})

	RegisterEndpoints(app)

	return app
}
