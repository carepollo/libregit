package middlewares

import (
	"github.com/carepollo/librecode/db"
	"github.com/carepollo/librecode/models"
	"github.com/gofiber/fiber/v2"
)

// guard to reject requests under protected routes when the user is not logged in.
// Checks that the current client has an active session.
func IsLogged(ctx *fiber.Ctx) error {
	_, err := db.GetUserSession(ctx.IP())
	if err != nil {
		return ctx.Redirect("/")
	}

	context, ok := ctx.Locals("globalData").(models.ContextData)
	if !ok {
		return ctx.Redirect("/")
	}

	if !context.IsLogged {
		return ctx.Redirect("/")
	}

	return ctx.Next()
}
