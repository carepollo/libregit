package controllers

import (
	"fmt"
	"log"
	"os"

	"github.com/carepollo/librecode/db"
	"github.com/carepollo/librecode/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleRegister(ctx *fiber.Ctx) error {
	username := ctx.FormValue("username")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")
	retypePassword := ctx.FormValue("confirmPassword")
	picture := ""

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

	// validate that user does not exist
	if db.NameIsRegistered(username) {
		return ctx.SendString("user " + username + " already exists ")
	}
	if db.EmailIsRegistered(email) {
		return ctx.SendString("email " + email + " already exists ")
	}

	// TODO: creating user's folder, creating personal repo for README

	// generate a random profile picture for user and store it in the user's directory
	filepath := fmt.Sprintf("%v/%v", utils.GlobalEnv.GitRoot, username)

	if err := os.MkdirAll(filepath, os.ModePerm); err != nil {
		log.Println("didn't add random pfp, fallback to default: ", err.Error())
	}

	picture, err := utils.GetRandomPfp(username, username)
	if err != nil {
		log.Println("didn't add random pfp, fallback to default: ", err.Error())
		picture = "/assets/defaultpfp.webp"
	}

	// hashing password
	password, err = utils.HashAndSalt(password)
	if err != nil {
		log.Println("didn't encrypted user password: " + err.Error())
		return ctx.Redirect("/register", fiber.StatusInternalServerError)
	}

	// creating user in database
	if err := db.CreateUser(username, email, password, picture); err != nil {
		log.Println("didn't create user on database: " + err.Error())
		return ctx.Redirect("/register", fiber.StatusInternalServerError)
	}

	mailMessage := fmt.Sprintf(`
	<h1>Welcome to LibreCode!</h1>
	<h4>We are happy that you chose us as your git provider</h4>
	<p>
		You have successfully created your <b>awesome</b> new account on LibreCode, to be able to log in you
		just need to <a href="%v/verify?account=%v">verify</a> your account.
	</p>
	`, utils.GlobalEnv.URLs.Project, username)
	err = utils.SendEmail([]string{email}, "Account verification", mailMessage)
	if err != nil {
		log.Println("couln't send mail: " + err.Error())
		return ctx.Redirect("/register", fiber.StatusInternalServerError)
	}

	// redirect to home where the feed should appear
	return ctx.Redirect("/login", fiber.StatusOK)
}
