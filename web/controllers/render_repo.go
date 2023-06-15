package controllers

import "github.com/gofiber/fiber/v2"

// handler to render the homepage of the repo
func RenderRepo(ctx *fiber.Ctx) error {
	return ctx.Render("repo", nil, "main")
}
