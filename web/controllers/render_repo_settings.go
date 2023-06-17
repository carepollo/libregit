package controllers

import "github.com/gofiber/fiber/v2"

func RenderRepoSettings(ctx *fiber.Ctx) error {
	return ctx.Render("views/user/repo_settings", nil, "main")
}
