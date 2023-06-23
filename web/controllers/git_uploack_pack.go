package controllers

import (
	"bytes"
	"fmt"
	"log"

	"github.com/carepollo/librecode/git"
	"github.com/carepollo/librecode/utils"
	"github.com/go-git/go-git/v5/plumbing/protocol/packp"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/gofiber/fiber/v2"
)

// handler for git-fetch, git-pull, git-clone, implementation of smart http
func HttpGitUploadPack(ctx *fiber.Ctx) error {
	ctx.Request().Header.Set("Content-Type", "application/x-git-uploack-pack-result")
	uploadRequest := packp.NewUploadPackRequest()
	body := bytes.NewReader(ctx.Request().Body())

	err := uploadRequest.Decode(body)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "could not decode body reader")
	}

	path := fmt.Sprintf("%v/%v/%v", utils.GlobalEnv.GitRoot, ctx.Params("user"), ctx.Params("repo"))
	endpoint, err := transport.NewEndpoint(path)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusNotFound, "could not find endpoint")
	}

	session, err := git.Gitserver.NewUploadPackSession(endpoint, nil)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "could not start upload pack session")
	}

	response, err := session.UploadPack(ctx.Context(), uploadRequest)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "could not upload pack")
	}

	err = response.Encode(ctx.Response().BodyWriter())
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "could not write the response")
	}

	return nil
}
