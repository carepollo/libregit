package controllers

import "github.com/gofiber/fiber/v2"

func RenderAccountSettings(ctx *fiber.Ctx) error {
	return ctx.Render("views/settings/account", nil, "main")
}
