package controller

import (
	"github.com/firdavs9512/freemail/services/freemail/databases"
	"github.com/kataras/iris/v12"
)

func ContactEmail(ctx iris.Context) {
	fullname := ctx.PostValue("name")
	email := ctx.PostValue("email")
	message := ctx.PostValue("message")

	_ = databases.CreateContactMessage(fullname, email, message)

	// ctx.SetCookieKV("message", "Message success send for admin!")
	// ctx.Redirect("/contact")
	ctx.HTML("Message success send for admin!")
}

// Get one email
func AdminGetOneEmail(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")

	email := databases.GetOneEmail(id)

	if email.ID == 0 {
		ctx.StatusCode(iris.StatusNotFound)
		return
	}

	ctx.ViewData("email", email)
	ctx.View("/admin/email-info.html")
}

// Delete email for base
func AdminDeleteEmail(ctx iris.Context) {
	id, _ := ctx.PostValueUint("id")

	err := databases.DeleteContactEmail(id)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.SetCookieKV("message", "Email delete successfull!")
}
