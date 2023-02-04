package controller

import (
	"github.com/firdavs9512/freemail/services/freemail/components"
	"github.com/firdavs9512/freemail/services/freemail/databases"
	"github.com/kataras/iris/v12"
)

func AdminLogin(ctx iris.Context) {
	ctx.ViewData("error", ctx.GetCookie("error"))
	ctx.RemoveCookie("error")
	ctx.View("/admin/auth-login.html")
}

func AdminLoginRequest(ctx iris.Context) {
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")

	admin := databases.AdminLogin(username, password)

	// if err != nil {
	// 	ctx.SetCookieKV("error", err.Error())
	// 	ctx.Redirect("/admin/login")
	// 	return
	// }

	if admin.Username == "" {
		ctx.SetCookieKV("error", "Login or Password error!")
		ctx.Redirect("/admin/login")
		return
	}

	session := components.Sess.Start(ctx)
	session.Set("admin", true)
	session.Set("username", admin.Username)

	ctx.SetCookieKV("message", "Welcome the FreeMail admin panel")
	ctx.Redirect("/admin")
}

func AdminMiddlware(ctx iris.Context) {
	auth, err := components.Sess.Start(ctx).GetBoolean("admin")
	if err != nil {
		ctx.SetCookieKV("error", "Admin not authenticated!")
		ctx.Redirect("/admin/login")
	}

	if !auth {
		ctx.SetCookieKV("error", "Admin not authenticated!")
		ctx.Redirect("/admin/login")
	}
	username := components.Sess.Start(ctx).GetString("username")

	admin := databases.AdminLoginForName(username)

	if admin.Username == "" {
		ctx.SetCookieKV("error", "Admin not authenticated!")
		ctx.Redirect("/admin/login")
	}

	ctx.Next()
}

func AdminLogout(ctx iris.Context) {
	session := components.Sess.Start(ctx)

	session.Delete("admin")
	session.Delete("username")
	ctx.Redirect("/")
}

// New admin create function
func AdminCreateAdminReq(ctx iris.Context) {
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	active, _ := ctx.PostValueBool("active")

	err := databases.AdminCreate(username, password, active)
	if err != nil {
		ctx.SetCookieKV("error", "Error for create admin!")
		return
	}

	ctx.SetCookieKV("message", "New admin created successfull!")
	ctx.Redirect("/admin")
}

// Admin status update reques
func AdminUpdateStatus(ctx iris.Context) {
	id, _ := ctx.PostValueUint("id")
	status, _ := ctx.PostValueBool("active")

	err := databases.AdminUpdate(id, status)

	if err != nil {
		ctx.SetCookieKV("error", "Error for update status admin!")
		return
	}

	ctx.SetCookieKV("message", "Admin status success updated!")

}

// Admin delete function
func AdminDeleteAdmins(ctx iris.Context) {
	id, _ := ctx.PostValueUint("id")

	err := databases.AdminDelete(id)

	if err != nil {
		ctx.SetCookieKV("error", "Error for update status admin!")
		return
	}

	ctx.SetCookieKV("message", "Admin deleted successfull!")
}
