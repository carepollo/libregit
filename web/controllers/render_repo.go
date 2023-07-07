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

	repo, err := db.GetRepoByOwnerAndName(owner, reponame)
	if err != nil {
		ctx.Redirect("/404")
	}

	contextData, ok := ctx.Locals("globalData").(models.ContextData)
	if !ok {
		log.Println("failed to get context data while rendering repo home")
		return ctx.Redirect("")
	}

	if repo.IsPrivate {
		if !contextData.IsLogged {
			return ctx.Redirect("/404")
		}

		if contextData.User.Name != owner {
			return ctx.Redirect("/404")
		}
	}

	readme, _ := git.GetReadme(owner, reponame)
	contextData.Readme = readme
	contextData.VisitedRepo = repo
	return ctx.Render("views/user/repo", contextData, "main")
}
