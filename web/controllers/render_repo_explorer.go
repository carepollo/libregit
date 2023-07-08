package controllers

import (
	"log"
	"strings"

	"github.com/carepollo/librecode/db"
	"github.com/carepollo/librecode/git"
	"github.com/carepollo/librecode/models"
	"github.com/gofiber/fiber/v2"
)

func RenderRepoExplorer(ctx *fiber.Ctx) error {
	ownername := ctx.Params("user")
	reponame := strings.TrimSuffix(ctx.Params("repo"), ".git")
	contextData, ok := ctx.Locals("globalData").(models.ContextData)
	if !ok {
		log.Println("failed to get context data on repo content redering")
		return ctx.Redirect("")
	}

	originalUrlContent := strings.Split(ctx.OriginalURL(), "/")
	contentPath := strings.Join(originalUrlContent[5:], "/")
	content, err := git.GetRepoFile(ownername, reponame, contentPath)
	if err != nil {
		log.Println("resource is not file:", err)
		if err.Error() == "entry is not file" {
			dir, err := git.GetRepoDir(ownername, reponame, contentPath)
			if err != nil {
				log.Println("resource is not dir:", err)
				dir = []map[string]interface{}{}
			}

			contextData.VisitedRepoDir = dir
		}
	} else {
		contextData.Readme = content
	}

	repo, err := db.GetRepoByOwnerAndName(ownername, reponame)
	if err != nil {
		log.Println("not found repo", err)
		return ctx.Redirect("/404")
	}

	contextData.VisitedRepo = repo
	return ctx.Render("views/user/repo_explorer", contextData, "main")
}
