package controllers

import (
	"log"
	"strings"

	"github.com/carepollo/librecode/db"
	"github.com/carepollo/librecode/models"
	"github.com/gofiber/fiber/v2"
)

// handler to render the homepage of the repo
func RenderRepoHome(ctx *fiber.Ctx) error {
	owner := ctx.Params("user")
	repoName := ctx.Params("repo")
	if strings.HasSuffix(repoName, ".git") {
		repoName = strings.TrimSuffix(repoName, ".git")
	}

	repo, err := db.GetRepoByOwnerAndName(owner, repoName)
	if err != nil {
		ctx.Redirect("/404")
	}

	contextData, ok := ctx.Locals("globalData").(models.ContextData)
	if !ok {
		log.Println("failed to get context data")
		ctx.Redirect("")
	}

	contextData.VisitedRepo = repo
	return ctx.Render("views/user/repo", contextData, "main")
}
