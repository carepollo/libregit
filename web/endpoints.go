package web

import (
	"github.com/carepollo/librecode/web/controllers"
	"github.com/carepollo/librecode/web/middlewares"
	"github.com/gofiber/fiber/v2"
)

// add the routes to the application
func RegisterEndpoints(app *fiber.App) {
	app.Use(middlewares.Logger)

	// serve files
	app.Static("/assets", "./assets")

	// web views
	app.Get("", controllers.RenderHome)
	app.Get("/login", controllers.RenderLogin)
	// app.Get("/register")
	// app.Get("/explore")
	// app.Get("/new")
	// app.Get("/settings")

	repos := app.Group("/:user/:repo")

	// git operations
	repos.Get("/info/refs", controllers.HttpGitInfoRefs)
	repos.Post("/git-upload-pack", controllers.HttpGitUploadPack)
	repos.Post("/git-receive-pack", controllers.HttpGitReceivePack)
}
