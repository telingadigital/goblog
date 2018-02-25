package routes

import (
	// "github.com/kataras/iris"
	controller "github.com/telingadigital/goblog/controller/api"
	"github.com/telingadigital/goblog/bootstrap"
)

func Api(app *bootstrap.App) {
	// Grouping
	apiRoutes := app.Party("/api")
	apiRoutes.Get("/",controller.Index)
}