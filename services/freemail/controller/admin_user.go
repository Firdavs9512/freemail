package controller

import (
	"github.com/firdavs9512/freemail/services/freemail/components"
	"github.com/firdavs9512/freemail/services/freemail/databases"
	"github.com/kataras/iris/v12"
)

func AdminCreateUserReq(ctx iris.Context) {
	firsname := ctx.PostValue("firstname")
	lastname := ctx.PostValue("lastname")
	email := ctx.PostValue("email")
	password := ctx.PostValue("password")

	pass, err := components.HashPass(password)
	if err != nil {
		ctx.SetCookieKV("error", "Error for create user!")
		ctx.Redirect("/admin")
	}

	user, err := databases.UserCreate(email, pass, lastname, firsname)

	if err != nil {
		ctx.SetCookieKV("error", "Error for create user!")
		ctx.Redirect("/admin")
	}

	if user.Email == "" {
		ctx.SetCookieKV("error", "Error for create user!")
		ctx.Redirect("/admin")
	}

	ctx.SetCookieKV("message", "New user created successfull!")
	ctx.Redirect("/admin")
}

// All users page paginator function
func AdminGetUsersPaginator(ctx iris.Context) {
	// page,_ := ctx.Params().GetUint("id")

}

func AdminAllUsers(ctx iris.Context) {
	// Send datas for template
	ctx.ViewData("error", ctx.GetCookie("error"))
	ctx.ViewData("message", ctx.GetCookie("message"))
	ctx.RemoveCookie("error")
	ctx.RemoveCookie("message")

	users := databases.GetAllUsers()
	ctx.ViewData("users", users)

	ctx.View("/admin/allusers.html")
}

// Users plan update handler
func AdminUserUpdate(ctx iris.Context) {
	plan, _ := ctx.PostValueInt("plan")
	id, _ := ctx.PostValueUint("id")

	err := databases.UserPlanUpdate(id, plan)

	if err != nil {
		ctx.SetCookieKV("error", "Error for update plan user!")
		return
	}

	ctx.SetCookieKV("message", "User plan update successful!")
}

// Delete user function
func AdminDeleteUser(ctx iris.Context) {
	id, _ := ctx.PostValueUint("id")

	err := databases.UserDelete(id)

	if err != nil {
		ctx.SetCookieKV("error", "Error for delete user!")
		return
	}

	ctx.SetCookieKV("message", "User success deleted!")
}
