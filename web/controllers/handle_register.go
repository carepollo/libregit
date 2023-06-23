package controllers

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/carepollo/librecode/db"
	"github.com/carepollo/librecode/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func HandleRegister(ctx *fiber.Ctx) error {
	username := strings.TrimSpace(ctx.FormValue("username"))
	email := strings.TrimSpace(ctx.FormValue("email"))
	password := ctx.FormValue("password")
	retypePassword := ctx.FormValue("confirmPassword")
	picture := ""
	id := uuid.New().String()

	// validating user intpus
	if !utils.ValidateName(username) {
		return ctx.Redirect("/register")
	}
	if !utils.ValidateEmail(email) {
		return ctx.Redirect("/register")
	}
	if !utils.ValidatePassword(password) {
		return ctx.Redirect("/register")
	}
	if password != retypePassword {
		return ctx.Redirect("/register")
	}

	// validate that user does not exist by checking user and email separately
	if db.NameIsRegistered(username) {
		return ctx.Redirect("/register")
	}
	if db.EmailIsRegistered(email) {
		return ctx.Redirect("/register")
	}

	// TODO: creating user's folder, creating personal repo for README
	filepath := fmt.Sprintf("%v/%v", utils.GlobalEnv.GitRoot, username)
	if err := os.MkdirAll(filepath, os.ModePerm); err != nil {
		log.Println("didn't add random pfp, fallback to default: ", err.Error())
	}

	// generate a random profile picture for user and store it in the user's directory
	picture, err := utils.GetRandomPfp(username, username)
	if err != nil {
		log.Println("didn't add random pfp, fallback to default: ", err.Error())
		picture = "/assets/defaultpfp.webp"
	}

	// hashing password
	password, err = utils.HashAndSalt(password)
	if err != nil {
		log.Println("didn't encrypted user password: " + err.Error()) // FIXME: add an emergency fallback
		return ctx.Redirect("/register")
	}

	// creating user in database
	if err := db.CreateUser(id, username, email, password, picture); err != nil {
		log.Println("didn't create user on database: " + err.Error())
		return ctx.Redirect("/register")
	}

	mailMessage := fmt.Sprintf(`
	<h1>Welcome to LibreCode!</h1>
	<h4>We are happy that you chose us as your git provider</h4>
	<p>
		You are one step ahead of your awesome journey on your new account on LibreCode, 
		to be able to log in you just need to <a href="%v/verify?userid=%v">verify</a> your account.
	</p>
	`, utils.GlobalEnv.URLs.Project, id)
	err = utils.SendEmail([]string{email}, "Account verification", mailMessage)
	if err != nil {
		log.Println("couln't send mail: " + err.Error())
		return ctx.Redirect("/register")
	}

	// redirect to home where the feed should appear
	return ctx.Redirect("/login")
}
