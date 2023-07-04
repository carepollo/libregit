package web

import (
	"github.com/carepollo/librecode/web/controllers"
	"github.com/carepollo/librecode/web/middlewares"
	"github.com/gofiber/fiber/v2"
)

// add the routes to the application
func RegisterEndpoints(app *fiber.App) {
	app.Use(middlewares.Logger)
	app.Use(middlewares.AppendUserData)

	// serve files
	app.Static("/assets", "./assets")

	files := app.Group("/media/:user")
	files.Get("/picture/:filename", controllers.GetUserPfp)

	// web views and ui-triggered actions
	app.Get("", controllers.RenderHome)

	app.Get("/login", controllers.RenderLogin)
	app.Post("/login", controllers.HandleLogin)
	app.Get("/logout", controllers.HandleLogout)

	app.Get("/register", controllers.RenderRegister)
	app.Post("/register", controllers.HandleRegister)

	app.Get("/verify", controllers.HandleVerify)
	app.Use(middlewares.IsLogged).Get("/new", controllers.RenderNewRepo)
	app.Use(middlewares.IsLogged).Post("/new", controllers.HandleNewRepo)

	settings := app.Group("settings")
	settings.Use(middlewares.IsLogged)
	settings.Get("/account", controllers.RenderAccountSettings)
	settings.Post("/account", controllers.HandleUserUpdate)
	settings.Post("/account/picture", controllers.SetUserPfp)

	user := app.Group("/:user")
	user.Get("", controllers.RenderUser)

	repos := user.Group("/:repo")
	repos.Get("", controllers.RenderRepoHome)
	repos.Get("/settings", controllers.RenderRepoSettings)

	// git operations
	repos.Get("/info/refs", controllers.HttpGitInfoRefs)
	repos.Post("/git-upload-pack", controllers.HttpGitUploadPack)
	repos.Post("/git-receive-pack", controllers.HttpGitReceivePack)

	// 404, when none of the previous routes matches the request
	app.Get("/404", controllers.RenderNotFound)
}
