package routers

import "github.com/kataras/iris/v12"

// 404 no found page function
func NotFoundHandler(ctx iris.Context) {
	ctx.HTML("<h1>page no found!</h1>")
}
