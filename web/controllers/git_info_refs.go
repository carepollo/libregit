package controllers

import (
	"fmt"
	"log"

	"github.com/carepollo/librecode/git"
	"github.com/go-git/go-git/v5/plumbing/format/pktline"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/gofiber/fiber/v2"
)

// handler for git all basic git operations, is a previous necessary step.
// Part of the implementation of the smart http
func HttpGitInfoRefs(ctx *fiber.Ctx) error {
	service := ctx.Query("service")
	ctx.Response().Header.Set("Content-Type", fmt.Sprintf("application/x-%v-advertisement", service))

	var err error
	var session transport.Session

	path := fmt.Sprintf("%v/%v/%v", git.GitPath, ctx.Params("user"), ctx.Params("repo"))
	endpoint, err := transport.NewEndpoint(path)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "failed at detecting endpoint")
	}

	if service == "git-upload-pack" {
		session, err = git.Gitserver.NewUploadPackSession(endpoint, nil)
	} else if service == "git-receive-pack" {
		session, err = git.Gitserver.NewReceivePackSession(endpoint, nil)
	} else {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "service not recognized")
	}
	if err != nil {
		log.Println(err)
		message := fmt.Sprintf("could not start %s session", service)
		return fiber.NewError(fiber.StatusInternalServerError, message)
	}

	refereces, err := session.AdvertisedReferencesContext(ctx.Context())
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "failed at getting advertised references context")
	}

	refereces.Prefix = [][]byte{
		[]byte(fmt.Sprintf("# service=%v", service)),
		pktline.Flush,
	}
	err = refereces.Encode(ctx.Response().BodyWriter())
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "could not process the writer")
	}

	return nil
}
