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
	app.Post("/login", controllers.HandleLogin)

	app.Get("/register", controllers.RenderRegister)
	// app.Get("/explore", controllers.RenderExplore)
	app.Get("/new", controllers.RenderNewRepo)

	settings := app.Group("settings")
	settings.Get("/account", controllers.RenderAccountSettings)

	user := app.Group("/:user")
	user.Get("", controllers.RenderUser)

	repos := user.Group("/:repo")
	repos.Get("", controllers.RenderRepoHome)
	repos.Get("/settings", controllers.RenderRepoSettings)

	// git operations
	repos.Get("/info/refs", controllers.HttpGitInfoRefs)
	repos.Post("/git-upload-pack", controllers.HttpGitUploadPack)
	repos.Post("/git-receive-pack", controllers.HttpGitReceivePack)
}
