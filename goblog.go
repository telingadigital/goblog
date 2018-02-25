package main

import (
	"github.com/telingadigital/goblog/bootstrap"
	"github.com/telingadigital/goblog/routes"
)

func RunApp() *bootstrap.App {
	app := bootstrap.New()
	app.NewApp()
	app.Configure(routes.Web, routes.Api)
	return app
}

func main() {
	app := RunApp()
	app.Listen()
}