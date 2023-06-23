package controllers

import (
	"log"

	"github.com/carepollo/librecode/db"
	"github.com/gofiber/fiber/v2"
)

func HandleLogout(ctx *fiber.Ctx) error {
	if err := db.DeleteUserSession(ctx.IP()); err != nil {
		log.Println(err)
	}
	return ctx.Redirect("/")
}
