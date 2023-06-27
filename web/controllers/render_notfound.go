package controllers

import "github.com/gofiber/fiber/v2"

func RenderNotFound(ctx *fiber.Ctx) error {
	return ctx.Render("404", nil, "main")
}
