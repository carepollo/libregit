package main

import (
	"github.com/carepollo/librecode/git"
	"github.com/carepollo/librecode/web"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	git.GitProjectRoot = "maiki"
	if err := godotenv.Load(); err != nil {
		panic("couldn't load environment variables file")
	}

	app := fiber.New()
	web.RegisterEndpoints(app)

	if err := app.Listen(":8080"); err != nil {
		panic("couldn't start server")
	}
}
