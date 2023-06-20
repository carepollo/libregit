package controllers

import (
	"fmt"
	"log"
	"os"

	"github.com/carepollo/librecode/db"
	"github.com/carepollo/librecode/git"
	"github.com/carepollo/librecode/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleRegister(ctx *fiber.Ctx) error {
	username := ctx.FormValue("username")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")
	retypePassword := ctx.FormValue("confirmPassword")

	// validating user intpus
	if !utils.ValidateName(username) {
		return ctx.SendString("username " + username + " not valid")
		// return ctx.Redirect("/register", fiber.StatusBadRequest)
	}
	if !utils.ValidateEmail(email) {
		return ctx.SendString("email " + email + " not valid")
	}
	if !utils.ValidatePassword(password) {
		return ctx.SendString("password " + password + " not valid")
	}
	if password != retypePassword {
		return ctx.SendString("password mismatch: " + password + " and " + retypePassword)
	}

	// creating user's folder, creating personal repo for README
	path := fmt.Sprintf("%v/%v/%v.git", git.GitPath, username, username)
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Println(err.Error())
		return ctx.Redirect("/register", fiber.StatusInternalServerError)
	}

	if _, err := git.CreateRepo(path); err != nil {
		log.Println(err.Error())
		return ctx.Redirect("/register", fiber.StatusInternalServerError)
	}

	// TODO: add README.md file to repo with sample text and commit it

	// hashing password
	password, err := utils.HashAndSalt(password)
	if err != nil {
		log.Println(err.Error())
		return ctx.Redirect("/register", fiber.StatusInternalServerError)
	}

	// creating user in database
	if err := db.CreateUser(username, email, password); err != nil {
		log.Println(err.Error())
		return ctx.Redirect("/register", fiber.StatusInternalServerError)
	}

	// redirect to home where the feed should appear
	return ctx.Redirect("/", fiber.StatusOK)
}
