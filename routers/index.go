package routers

import (
	"github.com/firdavs9512/freemail/services/freemail/controller"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

// Home login page controller
func Login(ctx iris.Context) {
	check := CheckLogin(ctx)
	if check {
		ctx.Redirect("/dashboard")
	}

	ctx.ViewData("error", ctx.GetCookie("error"))
	ctx.RemoveCookie("error")
	ctx.View("dashboard/auth-login.html")
}

// Home register page controller
func Register(ctx iris.Context) {

	check := CheckLogin(ctx)
	if check {
		ctx.Redirect("/dashboard")
	}

	ctx.ViewData("error", ctx.GetCookie("error"))
	ctx.RemoveCookie("error")
	ctx.View("dashboard/auth-register.html")
}

// Index page routers
func IndexRouter(app router.Party) {
	{
		app.Get("/", controller.Index).Name = "home"
		app.Get("/about", controller.About)
		app.Get("/pricing", controller.Pricing)
		app.Get("/privacy", controller.Privacy)
		app.Get("/contact", controller.Contact)
		app.Get("/signin", Login)
		app.Get("/signup", Register)
		app.Get("/reset-password", controller.ResetPassword)
		app.Get("/forgot-password", controller.ForgotPassword)
		app.Get("/logout", Logout).Name = "logout"
		app.Post("/auth-login", AuthLogin)
		app.Post("/auth-register", AuthRegister)
		app.Get("/services", controller.Services)
		app.Get("/service-promotion", controller.ServicesPromotion)
		app.Get("/reseller", controller.Reseller)

	}
}
