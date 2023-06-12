package main

import (
	"github.com/carepollo/librecode/utils"
	"github.com/carepollo/librecode/web"
)

func main() {
	utils.LoadEnv()

	app := web.CreateApp()
	app.Listen(":8080")
}
