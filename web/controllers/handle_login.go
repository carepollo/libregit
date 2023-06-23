package controllers

import (
	"log"
	"strings"

	"github.com/carepollo/librecode/db"
	"github.com/carepollo/librecode/models"
	"github.com/carepollo/librecode/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleLogin(ctx *fiber.Ctx) error {
	usermail := strings.TrimSpace(ctx.FormValue("useremail"))
	password := ctx.FormValue("password")

	user, err := db.GetUserByNameOrEmail(usermail)
	if err != nil {
		return ctx.Redirect("/login", fiber.StatusNotFound)
	}

	if !utils.CheckPassword(password, user.Password) {
		return ctx.Redirect("/login", fiber.StatusNotFound)
	}

	if user.Status != models.ACTIVE {
		return ctx.Redirect("/login", fiber.StatusNotAcceptable)
	}

	if err := db.SetUserSession(ctx.IP(), user); err != nil {
		log.Println(err)
		return ctx.Redirect("/login", fiber.StatusInternalServerError)
	}

	return ctx.Redirect("/")
}
