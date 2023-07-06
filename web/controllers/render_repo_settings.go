package controllers

import (
	"log"

	"github.com/carepollo/librecode/db"
	"github.com/carepollo/librecode/models"
	"github.com/gofiber/fiber/v2"
)

func RenderRepoSettings(ctx *fiber.Ctx) error {
	username := ctx.Params("user")
	reponame := ctx.Params("repo")
	contextData, ok := ctx.Locals("globalData").(models.ContextData)
	if !ok {
		log.Println("could not get context data on repo_settings render")
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	repo, err := db.GetRepoByOwnerAndName(username, reponame)
	if err != nil {
		log.Println("failed fetching repo data, return to homesite of repo")
		return ctx.Redirect(username + "/" + reponame)
	}

	contextData.VisitedRepo = repo
	contextData.ActiveTab = "settings"
	return ctx.Render("views/user/repo_settings", contextData, "main")
}
