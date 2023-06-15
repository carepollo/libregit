package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func RenderHome(ctx *fiber.Ctx) error {
	return ctx.Render("home", nil, "main")
}
