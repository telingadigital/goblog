package routes

import (
	// "github.com/kataras/iris"
	"github.com/telingadigital/goblog/controller"
	"github.com/telingadigital/goblog/bootstrap"
)

func Web(app *bootstrap.App) {

	app.Get("/",controller.Index)
}