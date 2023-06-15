package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// handler to render the home page, it covers the default landing page and the feed view when logged in
func RenderHome(ctx *fiber.Ctx) error {
	return ctx.Render("home", nil, "main")
}
