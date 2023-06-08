package web

import (
	"github.com/carepollo/librecode/web/controllers"
	"github.com/carepollo/librecode/web/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterEndpoints(app *fiber.App) {
	app.Use(middlewares.Logger)

	repos := app.Group("/:user/:repo")

	repos.Get("/info/refs", controllers.HttpGitInfoRefs)
	repos.Post("/git-upload-pack", controllers.HttpGitUploadPack)
	repos.Post("/git-receive-pack", controllers.HttpGitReceivePack)
}
