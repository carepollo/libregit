package web

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func RegisterEndpoints(app *fiber.App) {
	app.Get(":user/:repo/*", func(c *fiber.Ctx) error {
		path := c.Path()
		c.Params("user")
		direction := strings.Split(path, "/")

		if direction[2] == "git-upload-pack" {
			return nil
		}

		if direction[2] == "git-receive-pack" {
			return nil
		}

		c.Response().Header.Set("WWW-Authenticate", "Basic realm")
		return c.SendString("A")
	})
}
