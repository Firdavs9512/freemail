package controller

import (
	"github.com/firdavs9512/freemail/services/freemail/components"
	"github.com/firdavs9512/freemail/services/freemail/databases"
	"github.com/kataras/iris/v12"
)

func DashboardIndex(ctx iris.Context) {

	// Get user data for base
	email := components.Sess.Start(ctx).GetString("email")
	user := databases.UserLogin(email)
	ctx.ViewData("user", user)

	tokens := databases.GetFullTokens(user.ID)
	ctx.ViewData("tokens", tokens)

	// all requests
	sum := 0
	for i := 0; i < len(tokens); i++ {
		sum += int(tokens[i].Request)
	}
	ctx.ViewData("requests", sum)

	ctx.ViewData("templates", len(databases.GetTemplates()))

	ctx.ViewData("message", ctx.GetCookie("message"))
	ctx.ViewData("error", ctx.GetCookie("error"))

	ctx.RemoveCookie("error")
	ctx.RemoveCookie("message")
	ctx.View("/dashboard/index.html")
}

func DashboardCreateToken(ctx iris.Context) {
	url := components.ParseUrl(ctx)
	ctx.ViewData("url", url)

	// Get user data for base
	email := components.Sess.Start(ctx).GetString("email")
	user := databases.UserLogin(email)
	ctx.ViewData("user", user)

	// send all templates
	tmpl := databases.GetTemplates()
	ctx.ViewData("templates", tmpl)
	ctx.View("/dashboard/blank.html")
}

func DashboardProfile(ctx iris.Context) {
	// Get user data for base
	email := components.Sess.Start(ctx).GetString("email")
	user := databases.UserLogin(email)
	ctx.ViewData("user", user)

	url := components.ParseUrl(ctx)
	ctx.ViewData("url", url)
	ctx.View("/dashboard/profile.html")
}

func DashboardAllTemplates(ctx iris.Context) {
	// Get user data for base
	email := components.Sess.Start(ctx).GetString("email")
	user := databases.UserLogin(email)
	ctx.ViewData("user", user)

	temp := databases.GetTemplates()

	ctx.ViewData("templates", temp)
	ctx.View("/dashboard/card.html")
}
