package controllers

import (
	"log"

	"github.com/carepollo/librecode/db"
	"github.com/carepollo/librecode/git"
	"github.com/carepollo/librecode/models"
	"github.com/gofiber/fiber/v2"
)

// handler to render homepage of a user
func RenderUser(ctx *fiber.Ctx) error {
	var homeView string
	visitedUsername := ctx.Params("user")

	contextData, ok := ctx.Locals("globalData").(models.ContextData)
	if !ok {
		log.Println("failed retrieving context data")
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	tab := ctx.Query("tab")
	if tab == "repositories" {
		// if is tab of repos, get repos data
		homeView = "views/user/user_repos"
		includePrivate := visitedUsername == contextData.User.Name
		userRepos, err := db.GetReposByOwner(visitedUsername, includePrivate)
		if err != nil {
			log.Println("failed at getting user repos", err)
			userRepos = []models.Repo{}
		}

		contextData.VisitedUserRepos = userRepos
	} else {
		// if is main view, get readme of user
		homeView = "views/user/user"
		readme, err := git.GetReadme(visitedUsername, visitedUsername)
		if err != nil {
			log.Println("could not fetch user profile readme:", err)
		}

		contextData.Readme = readme
	}

	visited, err := db.GetUserByName(visitedUsername)
	if err != nil {
		return ctx.Redirect("/404")
	}
	contextData.VisitedUser = visited
	contextData.ActiveTab = tab

	return ctx.Render(homeView, contextData, "main")
}
