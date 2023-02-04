package routers

import (
	"github.com/firdavs9512/freemail/services/freemail/controller"
	"github.com/kataras/iris/v12/core/router"
)

// {{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
func DashboardRouter(app router.Party) {
	{
		app.Get("/", CheckLoginMiddleware, controller.DashboardIndex).Name = "dashboard"
		app.Get("/create-token", CheckLoginMiddleware, controller.DashboardCreateToken).Name = "createtoken"
		app.Get("/profile", CheckLoginMiddleware, controller.DashboardProfile).Name = "profile"
		app.Post("/create-token", CheckLoginMiddleware, controller.CreateTokenRouter)
		app.Post("/delete-token", CheckLoginMiddleware, controller.DeleteTokenForRequest).Name = "deletetoken"
		app.Get("/all-templates", controller.DashboardAllTemplates).Name = "alltemplates"
	}
}
