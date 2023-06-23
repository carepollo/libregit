package middlewares

import (
	"github.com/carepollo/librecode/db"
	"github.com/gofiber/fiber/v2"
)

// guard to reject requests under protected routes when the user is not logged in.
// Checks that the current client has an active session.
// TODO: add to the context the data of the user so the handler doesn't have to check the cache again.
func IsLogged(ctx *fiber.Ctx) error {
	_, err := db.GetUserSession(ctx.IP())
	if err != nil {
		return ctx.Redirect("/", fiber.StatusUnauthorized)
	}

	// ctx.Bind(fiber.Map{"userData", user})

	return ctx.Next()
}
