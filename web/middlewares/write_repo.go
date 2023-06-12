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
		return generateBasicAuthError(ctx)
	}

	prefix := "Basic "
	result, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return generateBasicAuthError(ctx)
	}

	userpassword := string(result)
	user, password, ok := strings.Cut(userpassword, ":")
	if !ok {
		return generateBasicAuthError(ctx)
	}

	if user != "admin" || password != "password" {
		return generateBasicAuthError(ctx)
	}

	return ctx.Next()
}

// shorthand to generate an unathorized error on basic auth middleware
func generateBasicAuthError(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("WWW-Authenticate", `Basic realm="librecode basic auth"`)
	return fiber.NewError(fiber.StatusUnauthorized)
}
