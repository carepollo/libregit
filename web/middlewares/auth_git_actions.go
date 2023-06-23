package middlewares

import (
	"encoding/base64"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ask for credentials when making git actions such as read/write operation that
// depends on the user permissions over repo.
func AuthGitActions(ctx *fiber.Ctx) error {
	// shorthand to generate an unathorized error on basic auth middleware
	basicAuthError := func(ctx *fiber.Ctx) error {
		ctx.Response().Header.Set("WWW-Authenticate", `Basic realm="librecode basic auth"`)
		return fiber.NewError(fiber.StatusUnauthorized)
	}

	auth := ctx.Get("Authorization")
	if auth == "" {
		return basicAuthError(ctx)
	}

	prefix := "Basic "
	result, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return basicAuthError(ctx)
	}

	userpassword := string(result)
	user, password, ok := strings.Cut(userpassword, ":")
	if !ok {
		return basicAuthError(ctx)
	}

	if user != "admin" || password != "password" {
		return basicAuthError(ctx)
	}

	return ctx.Next()
}
