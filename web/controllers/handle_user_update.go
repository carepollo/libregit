package controllers

import (
	"log"

	"github.com/carepollo/librecode/db"
	"github.com/gofiber/fiber/v2"
)

// handler to update the displayname and bio of user
func HandleUserUpdate(ctx *fiber.Ctx) error {
	displayName := ctx.FormValue("displayName")
	bio := ctx.FormValue("bio")

	user, err := db.GetUserSession(ctx.IP())
	if err != nil {
		log.Println("could not get session data:", err.Error())
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	user.Bio = bio
	user.DisplayName = displayName
	err = db.UpdateUser(user)
	if err != nil {
		log.Println("couldn't update user record:", err.Error())
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	err = db.SetUserSession(ctx.IP(), user)
	if err != nil {
		log.Println("couldn't update session:", err.Error())
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return ctx.Redirect("/settings/account")
}
