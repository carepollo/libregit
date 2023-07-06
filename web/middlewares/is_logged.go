package middlewares

import (
	"log"

	"github.com/carepollo/librecode/db"
	"github.com/gofiber/fiber/v2"
)

// guard to reject requests under protected routes when the user is not logged in.
// Checks that the current client has an active session.
func IsLogged(ctx *fiber.Ctx) error {
	_, err := db.GetUserSession(ctx.IP())
	if err != nil {
		log.Println("user is not logged:", err)
		return ctx.Redirect("/")
	}

	return ctx.Next()
}
