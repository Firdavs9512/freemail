package routers

import (
	"github.com/firdavs9512/freemail/services/freemail/controller"
	"github.com/kataras/iris/v12/core/router"
)

func FreeMailApiV1(app router.Party) {
	{
		app.Get("/test", controller.ApiTestV1)
		app.Get("/", controller.ApiForUsers)
		app.Post("/admin", controller.AdminSendEmailApi)
	}
}
