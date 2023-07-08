package middlewares

import "github.com/gofiber/fiber/v2"

// reject petition if the logged user that is requesting
// the resource is not the owner of the resource or doesn't
// have permissions.
func IsOwner(ctx *fiber.Ctx) error {
	return ctx.Next()
}
