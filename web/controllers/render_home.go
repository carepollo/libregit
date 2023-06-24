package controllers

import (
	"github.com/carepollo/librecode/db"
	"github.com/gofiber/fiber/v2"
)

// handler to render the home page, it covers the default landing page and the feed view when logged in
func RenderHome(ctx *fiber.Ctx) error {
	user, err := db.GetUserSession(ctx.IP())
	contextData := fiber.Map{
		"IsLogged": err == nil,
		"User":     user,
	}

	return ctx.Render("home", contextData, "main")
}
