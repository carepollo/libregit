package controllers

import (
	"github.com/carepollo/librecode/db"
	"github.com/gofiber/fiber/v2"
)

// verify account of user
func HandleVerify(ctx *fiber.Ctx) error {
	id := ctx.Query("userid")
	err := db.ActivateUser(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.Redirect("/")
}
