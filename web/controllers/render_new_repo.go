package controllers

import "github.com/gofiber/fiber/v2"

func RenderNewRepo(ctx *fiber.Ctx) error {
	return ctx.Render("new", nil, "main")
}
