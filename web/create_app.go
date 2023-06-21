// starting point of the application where all necessary configurations are being done
package web

import (
	"github.com/carepollo/librecode/db"
	"github.com/carepollo/librecode/git"
	"github.com/carepollo/librecode/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

// initialize web server with correspondant configuration and all application configurations
// assign to package variables some environment variables
func CreateApp() *fiber.App {
	git.GitPath = utils.GlobalEnv.GitRoot
	db.Open(utils.GlobalEnv.Storage.Db.Connection)

	// defining html rendering engine and configuring its behaviour
	engine := html.New("./templates", ".html")
	engine.Reload(utils.GlobalEnv.Production)

	// create fiber app and append endpoints
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		Views:         engine,
	})

	RegisterEndpoints(app)

	return app
}
