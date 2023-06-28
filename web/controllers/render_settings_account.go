package controllers

import (
	"log"

	"github.com/carepollo/librecode/models"
	"github.com/gofiber/fiber/v2"
)

func RenderAccountSettings(ctx *fiber.Ctx) error {
	contextData, ok := ctx.Locals("globalData").(models.ContextData)
	if !ok {
		log.Println("not found contextData")
		return ctx.Redirect("/")
	}

	return ctx.Render("views/settings/account", contextData, "main")
}
