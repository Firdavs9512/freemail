package components

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

var (
	cookieNameForSessionID = "_cookie"
	Sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)

func GetSessionForData(ctx iris.Context, data string) string {
	res := Sess.Start(ctx).GetString(data)
	return res
}
