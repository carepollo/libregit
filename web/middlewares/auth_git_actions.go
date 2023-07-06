package middlewares

import (
	"encoding/base64"
	"log"
	"strings"

	"github.com/carepollo/librecode/db"
	"github.com/carepollo/librecode/utils"
	"github.com/gofiber/fiber/v2"
)

// a way to transfer as an object the user password passed to the auth method,
// it contains only the user name and the user password
type credentials struct {
	user     string
	password string
}

// ask for credentials when making git actions such as read/write operation that
// depends on the user permissions over repo.
// works only under the /user/repo.git/info/refs endpoint
func AuthGitActions(ctx *fiber.Ctx) error {
	username := ctx.Params("user")
	reponame := strings.TrimSuffix(ctx.Params("repo"), ".git")
	service := strings.TrimPrefix(ctx.Query("service"), "git-")
	credentials := credentials{}
	repo, err := db.GetRepoByOwnerAndName(username, reponame)
	if err != nil {
		log.Println("repo requested not found on database", err)
		return fiber.NewError(fiber.StatusNotFound)
	}

	user, err := db.GetUserByName(username)
	if err != nil {
		log.Println("owner of repo requested not found on database", err)
		return fiber.NewError(fiber.StatusNotFound)
	}

	credentials.user = user.Name
	credentials.password = user.Password

	// if request wants to push to a private repo
	if service == "receive-pack" {
		return triggerBasicAuth(ctx, credentials)
	}

	// if repo is private, any action shall be authenticate
	if repo.IsPrivate {
		return triggerBasicAuth(ctx, credentials)
	}

	return ctx.Next()
}

// TODO: implement basic and bearer auth in an extensible way

func triggerBasicAuth(ctx *fiber.Ctx, creds credentials) error {
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

	if user != creds.user {
		log.Println("user provided is not valid")
		return basicAuthError(ctx)
	}

	if !utils.CheckPassword(password, creds.password) {
		log.Println("password provided is not valid")
		return basicAuthError(ctx)
	}

	return ctx.Next()
}

// func triggerBearerAuth(ctx *fiber.Ctx) error { return ctx.Next() }
