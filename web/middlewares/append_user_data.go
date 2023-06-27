package middlewares

import (
	"github.com/carepollo/librecode/db"
	"github.com/carepollo/librecode/models"
	"github.com/gofiber/fiber/v2"
)

// add to the context of the request data of the user if its logged in so it can be used in
// the rendered template or as variable for the handler
func AppendUserData(ctx *fiber.Ctx) error {
	user, err := db.GetUserSession(ctx.IP())
	contextData := models.ContextData{
		IsLogged: err == nil,
		User:     user,
	}
	ctx.Locals("globalData", contextData)

	return ctx.Next()
}
