package routers

import (
	"github.com/firdavs9512/freemail/services/freemail/controller"
	"github.com/kataras/iris/v12/core/router"
)

func DocsRouter(app router.Party) {
	{
		app.Get("/", controller.DocsIndex)
	}
}
