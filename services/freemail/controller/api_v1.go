package controller

import (
	"github.com/firdavs9512/freemail/services/freemail/components"
	"github.com/firdavs9512/freemail/services/freemail/databases"
	"github.com/kataras/iris/v12"
)

func ApiTestV1(ctx iris.Context) {
	email := ctx.URLParamDefault("email", "firdavs20fe@gmail.com")
	subject := ctx.URLParamDefault("subject", "Subject Test")
	message := ctx.URLParamDefault("message", "Message Test")

	err := components.SendEmail(email, message, subject, "")

	if err != nil {
		ctx.Writef("Error send test email: %s", err.Error())
	}

	ctx.Write([]byte("Ok"))
}

func ApiForUsers(ctx iris.Context) {
	url := ctx.URLParam("token")
	subject := ctx.URLParam("subject")
	message := ctx.URLParam("message")
	email := ctx.URLParam("email")

	token := databases.GetDataForUrl(url)

	if token.Url == "" {
		ctx.JSON(iris.Map{
			"message": "Your token not found!",
		})
		return
	}

	if token.Maxreq < token.Request {
		ctx.JSON(iris.Map{
			"message": "Token request max!",
		})
		return
	}

	err := databases.TokenRequestUpdate(token.ID, token.Request+1)
	if err != nil {
		ctx.JSON(iris.Map{
			"message": "Your token not found!",
		})
		return
	}

	// Send email
	res := components.SendEmail(email, message, token.Template, subject)

	if res != nil {
		ctx.JSON(iris.Map{
			"message": "Your token not found!",
		})
		return
	}

	ctx.JSON(iris.Map{
		"message": "Success",
	})

}

// Api for admin
func AdminSendEmailApi(ctx iris.Context) {
	email := ctx.PostValue("email")
	subject := ctx.PostValue("subject")
	message := ctx.PostValue("message")

	res := components.SendEmail(email, message, "template.html", subject)

	if res != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{
			"message": "Error send email",
		})
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{
		"message": "Ok",
	})
}
