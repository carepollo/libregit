package controllers

import (
	"path/filepath"

	"github.com/carepollo/librecode/utils"
	"github.com/gofiber/fiber/v2"
)

// handler to search user profile picture in server and return it.
func GetUserPfp(ctx *fiber.Ctx) error {
	path := filepath.Join(utils.GlobalEnv.GitRoot, ctx.Params("user"), ctx.Params("filename"))
	return ctx.SendFile(path)
}
