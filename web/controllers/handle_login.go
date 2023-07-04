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
		log.Println("not found user by email or email:", err)
		return ctx.Redirect("/login")
	}

	if !utils.CheckPassword(password, user.Password) {
		log.Println("the password doesn't match")
		return ctx.Redirect("/login")
	}

	if user.Status != models.ACTIVE {
		log.Println("the user given is not verified")
		return ctx.Redirect("/login", fiber.StatusNotAcceptable)
	}

	if err := db.SetUserSession(ctx.IP(), user); err != nil {
		log.Println("couldn't store user session", err)
		return ctx.Redirect("/login", fiber.StatusInternalServerError)
	}

	return ctx.Redirect("/")
}
