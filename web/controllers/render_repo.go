package controllers

import "github.com/gofiber/fiber/v2"

// handler to render the homepage of the repo
func RenderRepoHome(ctx *fiber.Ctx) error {
	return ctx.Render("views/user/repo", fiber.Map{"sample": sampleMD}, "main")
}
