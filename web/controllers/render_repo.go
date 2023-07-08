package controllers

import (
	"log"
	"strings"

	"github.com/carepollo/librecode/db"
	"github.com/carepollo/librecode/git"
	"github.com/carepollo/librecode/models"
	"github.com/gofiber/fiber/v2"
)

// handler to render the homepage of the repo
func RenderRepoHome(ctx *fiber.Ctx) error {
	owner := ctx.Params("user")
	reponame := strings.TrimSuffix(ctx.Params("repo"), ".git")

	// getting data of repo for validations
	repo, err := db.GetRepoByOwnerAndName(owner, reponame)
	if err != nil {
		ctx.Redirect("/404")
	}

	// get data shared across app
	contextData, ok := ctx.Locals("globalData").(models.ContextData)
	if !ok {
		log.Println("failed to get context data while rendering repo home")
		return ctx.Redirect("")
	}

	// prevent views from non-owner users
	if repo.IsPrivate {
		if !contextData.IsLogged {
			return ctx.Redirect("/404")
		}

		if contextData.User.Name != owner {
			return ctx.Redirect("/404")
		}
	}

	// getting the content of the README.md file if there is any
	readme, _ := git.GetReadme(owner, reponame)
	contextData.Readme = readme
	contextData.VisitedRepo = repo

	// getting repo content at root of the worktree
	dirs, err := git.GetRepoDir(owner, reponame, "")
	if err != nil {
		log.Println("failed to get repo dirs, falling back to default value:", err)
		dirs = []map[string]interface{}{}
	}

	contextData.VisitedRepoDir = dirs
	return ctx.Render("views/user/repo", contextData, "main")
}
