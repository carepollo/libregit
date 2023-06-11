package web

import (
	"github.com/carepollo/librecode/web/controllers"
	"github.com/carepollo/librecode/web/middlewares"
	"github.com/gofiber/fiber/v2"
)

// add the routes to the application
func RegisterEndpoints(app *fiber.App) {
	app.Use(middlewares.Logger)

	// web views
	// app.Get("")
	// app.Get("/login")
	// app.Get("/register")
	// app.Get("/explore")
	// app.Get("/new")
	// app.Get("/settings")

	repos := app.Group("/:user/:repo")

	// git operations
	repos.Use(middlewares.AuthWriteRepo).Get("/info/refs", controllers.HttpGitInfoRefs)
	repos.Post("/git-upload-pack", controllers.HttpGitUploadPack)
	repos.Post("/git-receive-pack", controllers.HttpGitReceivePack)
}
