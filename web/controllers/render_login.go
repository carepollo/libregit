package controllers

import "github.com/gofiber/fiber/v2"

// handler for rendering the html of the login page
func RenderLogin(ctx *fiber.Ctx) error {
	return ctx.Render("login", nil, "main")
}
