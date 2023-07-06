package controllers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/carepollo/librecode/db"
	"github.com/carepollo/librecode/git"
	"github.com/carepollo/librecode/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleNewRepo(ctx *fiber.Ctx) error {
	reponame := ctx.FormValue("reponame")
	visibility := ctx.FormValue("visibility") == "private"
	user, err := db.GetUserSession(ctx.IP())
	if err != nil {
		log.Println("not found user session data")
		return ctx.Redirect("/new")
	}

	path := filepath.Join(utils.GlobalEnv.GitRoot, user.Name, reponame+".git")
	_, err = git.CreateRepo(path)
	if err != nil {
		log.Println("failed to create repository:", err)
		return ctx.Redirect("/new")
	}

	err = db.CreateRepo(user.Name, reponame, visibility)
	if err != nil {
		log.Println("could not register repo:", err)
		err = os.Remove(path)
		if err != nil {
			log.Println("failed on cleaning failed repo:", err)
		}
		return ctx.Redirect("/new")
	}

	user.AmountRepositories++
	err = db.UpdateUser(user)
	if err != nil {
		log.Println("failed to update user data", err)
		return ctx.Redirect("/new")
	}

	newRepoPath := fmt.Sprintf("/%v/%v", user.Name, reponame)
	return ctx.Redirect(newRepoPath)
}
