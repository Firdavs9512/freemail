package controller

import (
	"github.com/firdavs9512/freemail/services/freemail/databases"
	"github.com/kataras/iris/v12"
)

func AdminDashboard(ctx iris.Context) {

	// Send datas for template
	ctx.ViewData("error", ctx.GetCookie("error"))
	ctx.ViewData("message", ctx.GetCookie("message"))
	ctx.RemoveCookie("error")
	ctx.RemoveCookie("message")

	ctx.ViewData("users", len(databases.GetAllUsers()))
	ctx.ViewData("tokens", len(databases.GetAllTokens()))
	ctx.ViewData("templates", len(databases.GetTemplates()))
	ctx.ViewData("primumusers", len(databases.GetPrimumUsers()))

	ctx.View("/admin/index.html")

}

func AdminPayments(ctx iris.Context) {
	// Send datas for template
	ctx.ViewData("error", ctx.GetCookie("error"))
	ctx.ViewData("message", ctx.GetCookie("message"))
	ctx.RemoveCookie("error")
	ctx.RemoveCookie("message")

	ctx.View("/admin/payments.html")
}

func AdminCreateTemplate(ctx iris.Context) {
	// Send datas for template
	ctx.ViewData("error", ctx.GetCookie("error"))
	ctx.ViewData("message", ctx.GetCookie("message"))
	ctx.RemoveCookie("error")
	ctx.RemoveCookie("message")

	ctx.View("/admin/Template.html")
}

func AdminAllTemplates(ctx iris.Context) {
	// Send datas for template
	ctx.ViewData("error", ctx.GetCookie("error"))
	ctx.ViewData("message", ctx.GetCookie("message"))
	ctx.RemoveCookie("error")
	ctx.RemoveCookie("message")

	ctx.ViewData("templates", GetAllTemplates())

	ctx.View("/admin/Alltemplates.html")
}

func AdminCreateUser(ctx iris.Context) {
	// Send datas for template
	ctx.ViewData("error", ctx.GetCookie("error"))
	ctx.ViewData("message", ctx.GetCookie("message"))
	ctx.RemoveCookie("error")
	ctx.RemoveCookie("message")

	ctx.View("/admin/create-user.html")
}

func AdminContactEmails(ctx iris.Context) {
	// Send datas for template
	ctx.ViewData("error", ctx.GetCookie("error"))
	ctx.ViewData("message", ctx.GetCookie("message"))
	ctx.RemoveCookie("error")
	ctx.RemoveCookie("message")

	all := databases.AllContactEmails()
	ctx.ViewData("emails", all)

	ctx.View("/admin/emails.html")
}

func AdminSendEmail(ctx iris.Context) {
	// Send datas for template
	ctx.ViewData("error", ctx.GetCookie("error"))
	ctx.ViewData("message", ctx.GetCookie("message"))
	ctx.RemoveCookie("error")
	ctx.RemoveCookie("message")

	ctx.View("/admin/email-send.html")
}

func AdminCreateAdmin(ctx iris.Context) {
	// Send datas for template
	ctx.ViewData("error", ctx.GetCookie("error"))
	ctx.ViewData("message", ctx.GetCookie("message"))
	ctx.RemoveCookie("error")
	ctx.RemoveCookie("message")

	ctx.View("/admin/create-admin.html")
}

func AdminAllAdmins(ctx iris.Context) {
	// Send datas for template
	ctx.ViewData("error", ctx.GetCookie("error"))
	ctx.ViewData("message", ctx.GetCookie("message"))
	ctx.RemoveCookie("error")
	ctx.RemoveCookie("message")

	// Send all admins variable
	ctx.ViewData("admins", databases.GetAllAdmins())

	ctx.View("/admin/alladmins.html")
}
