package middlewares

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Logger(ctx *fiber.Ctx) error {
	log.Printf("%s %s", ctx.Method(), ctx.Path())
	return ctx.Next()
}
