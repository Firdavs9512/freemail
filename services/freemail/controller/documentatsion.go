package controller

import "github.com/kataras/iris/v12"

func DocsIndex(ctx iris.Context) {
	ctx.View("idocs/index.html")
}
