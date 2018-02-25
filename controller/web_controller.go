package controller

import(
	"github.com/kataras/iris"
)

func Index(ctx iris.Context) {
	ctx.Writef("Hello World")
}