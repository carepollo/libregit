package controllers

import "github.com/gofiber/fiber/v2"

// handler to render the register new account view
func RenderRegister(ctx *fiber.Ctx) error {
	return ctx.Render("register", nil, "main")
}
