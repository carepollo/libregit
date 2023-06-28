package controllers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/carepollo/librecode/db"
	"github.com/carepollo/librecode/utils"
	"github.com/gofiber/fiber/v2"
)

// handler to change the profile picture of the user
func SetUserPfp(ctx *fiber.Ctx) error {
	// get file from form data
	file, err := ctx.FormFile("picture")
	if err != nil {
		log.Println("couldn't get local data", err.Error())
		return fiber.NewError(fiber.StatusNotAcceptable, err.Error())
	}

	// get user data from cache stored in request context from middleware
	user, err := db.GetUserSession(ctx.IP())
	if err != nil {
		log.Println("couldn't get local data", err.Error())
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// saving received file and keeping name of previous file
	previousUrl := strings.Split(user.Picture, "/")
	previousName := previousUrl[len(previousUrl)-1]
	path := filepath.Join(utils.GlobalEnv.GitRoot, user.Name, file.Filename)
	err = ctx.SaveFile(file, path)
	if err != nil {
		log.Println("couldn't get local data", err.Error())
		return fiber.NewError(fiber.StatusNotAcceptable, err.Error())
	}

	// updating record of user in DB
	user.Picture = fmt.Sprintf("%v/%v/%v/%v", utils.GlobalEnv.URLs.Project, "media", user.Name, file.Filename)
	err = db.UpdateUser(user)
	if err != nil {
		log.Println("couldn't get local data", err.Error())
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// update user session
	err = db.SetUserSession(ctx.IP(), user)
	if err != nil {
		log.Println("couldn't get local data", err.Error())
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// delete previous pfp from server disk
	err = os.Remove(filepath.Join(utils.GlobalEnv.GitRoot, user.Name, previousName))
	if err != nil {
		log.Println(err)
	}

	return ctx.Redirect("/settings/account")
}
