package controller

import (
	"github.com/firdavs9512/freemail/services/freemail/components"
	"github.com/firdavs9512/freemail/services/freemail/databases"
	"github.com/kataras/iris/v12"
)

func CreateTokenRouter(ctx iris.Context) {
	name := ctx.PostValue("name")
	check := ctx.PostValue("template")

	email := components.Sess.Start(ctx).GetString("email")

	token, err := databases.CreateToken(email, name, check)

	if err != nil {
		ctx.SetCookieKV("error", err.Error())
		ctx.Redirect("/dashboard")
		return
	}

	if token.Url == "" {
		ctx.SetCookieKV("error", "Error for create token!")
		ctx.Redirect("/dashboard")
		return
	}

	ctx.SetCookieKV("message", "Token created succsessfull !")
	ctx.Redirect("/dashboard")

}

func DeleteTokenForRequest(ctx iris.Context) {
	id, _ := ctx.PostValueInt("id")
	email := components.Sess.Start(ctx).GetString("email")

	user := databases.UserLogin(email)

	err := databases.TokenFullDelete(uint(id), user.ID)

	if err != nil {
		ctx.SetCookieKV("error", err.Error())
		return
	}

	ctx.SetCookieKV("message", "Token is deleted successfull!")

}
