package controllers

import "github.com/gofiber/fiber/v2"

func HandleLogin(ctx *fiber.Ctx) error {
	// ctx.Redirect("/")
	return ctx.SendString(ctx.FormValue("useremail") + ctx.FormValue("password"))
}
