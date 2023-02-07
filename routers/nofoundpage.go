package routers

import "github.com/kataras/iris/v12"

// 404 no found page function
func NotFoundHandler(ctx iris.Context) {
	ctx.View("home/404.html")
}
