package main

import (
	"github.com/carepollo/librecode/db"
	"github.com/carepollo/librecode/utils"
	"github.com/carepollo/librecode/web"
)

func main() {
	defer db.Close()

	utils.LoadEnv()

	app := web.CreateApp()
	app.Listen(":8080")
}
