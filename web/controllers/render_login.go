package controllers

import "github.com/gofiber/fiber/v2"

func RenderLogin(ctx *fiber.Ctx) error {
	return ctx.Render("login", nil, "main")
}
