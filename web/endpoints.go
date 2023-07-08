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

	// 404, when none of the previous routes matches the request
	app.Get("/404", controllers.RenderNotFound)

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
	app.Get("/new", middlewares.IsLogged, controllers.RenderNewRepo)
	app.Post("/new", middlewares.IsLogged, controllers.HandleNewRepo)

	settings := app.Group("settings", middlewares.IsLogged)
	settings.Get("/account", controllers.RenderAccountSettings)
	settings.Post("/account", controllers.HandleUserUpdate)
	settings.Post("/account/picture", controllers.SetUserPfp)

	user := app.Group("/:user")
	user.Get("", controllers.RenderUser)

	repos := user.Group("/:repo")
	repos.Get("", controllers.RenderRepoHome)
	repos.Get("/settings", middlewares.IsLogged, controllers.RenderRepoSettings)
	repos.Post("/settings", middlewares.IsLogged, controllers.HandleUpdateRepoSettings)
	repos.Get("/src/:branch/*", middlewares.IsLogged, controllers.RenderRepoExplorer)

	// git operations
	repos.Get("/info/refs", middlewares.AuthGitActions, controllers.HttpGitInfoRefs)
	repos.Post("/git-upload-pack", controllers.HttpGitUploadPack)
	repos.Post("/git-receive-pack", controllers.HttpGitReceivePack)
}
