package middlewares

import (
	"encoding/base64"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ask for credentials to write and check for write permissions over repo
func AuthWriteRepo(ctx *fiber.Ctx) error {
	auth := ctx.Get("Authorization")
	if auth == "" {
		return fiber.NewError(fiber.StatusUnauthorized)
	}

	prefix := "Basic "
	result, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	userpassword := string(result)
	user, password, ok := strings.Cut(userpassword, ":")
	if !ok {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	if user != "admin" || password != "password" {
		return fiber.NewError(fiber.StatusUnauthorized)
	}

	return ctx.Next()
}
