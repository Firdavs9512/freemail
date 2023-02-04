package routers

import (
	"github.com/firdavs9512/freemail/services/freemail/components"
	"github.com/firdavs9512/freemail/services/freemail/databases"
	"github.com/kataras/iris/v12"
)

// var (
// 	cookieNameForSessionID = "_cookie"
// 	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
// )

// User registration function
func AuthRegister(ctx iris.Context) {
	lastname := ctx.PostValue("last_name")
	firstname := ctx.PostValue("frist_name")
	email := ctx.PostValue("email")
	password := ctx.PostValue("password")
	agree := ctx.PostValue("agree")
	passwordconfirm := ctx.PostValue("password-confirm")

	if components.IsEmmty(lastname) || components.IsEmmty(firstname) {
		ctx.SetCookieKV("error", "Lastname or Firstname is emmty!")
		ctx.Redirect("/signup")
		return
	}
	if components.IsChecked(agree) {
		ctx.SetCookieKV("error", "I agree with the terms and conditions not checked")
		ctx.Redirect("/signup")
		return
	}
	if password != passwordconfirm {
		ctx.SetCookieKV("error", "Password and password confirm not equel!")
		ctx.Redirect("/signup")
		return
	}
	if len(password) < 7 {
		ctx.SetCookieKV("error", "Password must be longer than 8 characters!")
		ctx.Redirect("/signup")
		return
	}

	if !databases.UserControll(email) {
		ctx.SetCookieKV("error", "This user has already registered")
		ctx.Redirect("/signup")
		return
	}

	pass, err := components.HashPass(password)
	if err != nil {
		ctx.SetCookieKV("error", "Error for create user!")
		ctx.Redirect("/signup")
		return
	}
	user, err := databases.UserCreate(email, pass, lastname, firstname)

	if err != nil {
		ctx.SetCookieKV("error", "Error for create user!")
		ctx.Redirect("/signup")
		return
	}

	session := components.Sess.Start(ctx)
	session.Set("authenticated", true)
	session.Set("id", user.ID)
	session.Set("email", user.Email)

	ctx.SetCookieKV("message", "Welcome to FreeMail free api service!")
	ctx.Redirect("/dashboard")

}

// User middlevare function
func CheckLoginMiddleware(ctx iris.Context) {
	auth, err := components.Sess.Start(ctx).GetBoolean("authenticated")
	if err != nil {
		ctx.SetCookieKV("error", "User not authenticated!")
		ctx.Redirect("/signup")
	}

	if !auth {
		ctx.SetCookieKV("error", "User not authenticated!")
		ctx.Redirect("/signup")
	}
	email := components.Sess.Start(ctx).GetString("email")

	user := databases.UserLogin(email)

	if user.Email == "" {
		ctx.SetCookieKV("error", "User not authenticated!")
		ctx.Redirect("/signup")
	}

	ctx.Next()

}

// User middlevare function
func CheckLogin(ctx iris.Context) bool {
	auth, err := components.Sess.Start(ctx).GetBoolean("authenticated")
	if err != nil {
		return false
	}

	if !auth {
		return false
	}

	return true

}

// User logout function
func Logout(ctx iris.Context) {
	session := components.Sess.Start(ctx)

	// Revoke users authentication
	// session.Set("authenticated", false)
	// Or to remove the variable:

	session.Delete("authenticated")
	session.Delete("email")
	session.Delete("id")
	ctx.Redirect("/")
}

// User login function
func AuthLogin(ctx iris.Context) {
	email := ctx.PostValue("email")
	password := ctx.PostValue("password")

	user := databases.UserLogin(email)

	err := components.HashControll(user.Password, password)
	if err != nil {
		ctx.SetCookieKV("error", "Login or password not correct!")
		ctx.Redirect("/signin")
		return
	}

	if user.Email == "" {
		ctx.SetCookieKV("error", "Login or password not correct!")
		ctx.Redirect("/signin")
		return
	}

	session := components.Sess.Start(ctx)
	session.Set("authenticated", true)
	session.Set("id", user.ID)
	session.Set("email", user.Email)

	ctx.SetCookieKV("message", "Welcome to FreeMail free api service!")
	ctx.Redirect("/dashboard")
}
