package controllers

import (
	"github.com/carepollo/librecode/db"
	"github.com/carepollo/librecode/models"
	"github.com/gofiber/fiber/v2"
)

// handler to render homepage of a user
func RenderUser(ctx *fiber.Ctx) error {
	var homeView string
	tab := ctx.Query("tab")
	if tab == "repositories" {
		homeView = "views/user/list_repos"
	} else {
		homeView = "views/user/user"
	}

	contextData, ok := ctx.Locals("globalData").(models.ContextData)
	if !ok {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	visited, err := db.GetUserByName(ctx.Params("user"))
	if err != nil {
		return ctx.Redirect("/404")
	}
	contextData.VisitedUser = visited
	contextData.ActiveTab = tab

	return ctx.Render(homeView, contextData, "main")
}
