package controllers

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5/plumbing/protocol/packp"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/gofiber/fiber/v2"
)

// handler for git-push
func HttpGitReceivePack(ctx *fiber.Ctx) error {
	ctx.Request().Header.Set("Content-Type", "application/x-git-receive-pack-result")
	receiveRequest := packp.NewReferenceUpdateRequest()
	body := ctx.Request().BodyStream()
	err := receiveRequest.Decode(body)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "could not decode body reader")
	}
	ctx.Request().CloseBodyStream()

	path := fmt.Sprintf("%v/%v/%v", gitpath, ctx.Params("user"), ctx.Params("repo"))
	endpoint, err := transport.NewEndpoint(path)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusNotFound, "could not find endpoint")
	}

	session, err := gitserver.NewReceivePackSession(endpoint, nil)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "could not start receive pack session")
	}

	response, err := session.ReceivePack(ctx.Context(), receiveRequest)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "could not receive pack")
	}

	err = response.Encode(ctx.Response().BodyWriter())
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "could not write the response")
	}

	return nil
}
