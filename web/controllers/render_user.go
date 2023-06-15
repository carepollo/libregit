package controllers

import "github.com/gofiber/fiber/v2"

// handler to render homepage of a user
func RenderUser(ctx *fiber.Ctx) error {
	text :=
		`
	# this is a test
                        
	**lol** I would _love_ to to do that
	I would not do any of this without you [guys](https://google.com).
	I would like to share this pic.

	![img test](http://localhost:8080/assets/foss.png)

	and this code:
	`
	return ctx.Render("views/user/user", fiber.Map{
		"sample": text,
	}, "main")
}
