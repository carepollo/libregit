package controllers

import (
	"fmt"
	"log"

	"github.com/carepollo/librecode/db"
	"github.com/carepollo/librecode/git"
	"github.com/carepollo/librecode/utils"
	"github.com/gofiber/fiber/v2"
)

func HandleUpdateRepoSettings(ctx *fiber.Ctx) error {
	username := ctx.Params("user")
	reponame := ctx.Params("repo")
	newRepoName := ctx.FormValue("repo")
	description := ctx.FormValue("description")
	currentPath := "/" + username + "/" + reponame
	isPrivate := ctx.FormValue("visibility") == "private"

	repo, err := db.GetRepoByOwnerAndName(username, reponame)
	if err != nil {
		log.Println("failed at fething stored repo:", err)
		return ctx.Redirect(currentPath)
	}

	err = git.RenameRepo(reponame, newRepoName, username)
	if err != nil {
		log.Println("failed at renaming repo folder:", err)
		return ctx.Redirect(currentPath)
	}

	repo.Name = newRepoName
	repo.Location = fmt.Sprintf("%v/%v/%v", utils.GlobalEnv.URLs.Project, username, newRepoName)
	repo.IsPrivate = isPrivate
	repo.Description = description
	err = db.UpdateRepo(repo)
	if err != nil {
		log.Println("failed at updating stored repo data:", err)
		return ctx.Redirect(currentPath)
	}

	return ctx.Redirect(repo.Location)
}
