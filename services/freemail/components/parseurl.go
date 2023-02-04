package components

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func ParseUrl(ctx iris.Context) string {
	url := ctx.FullRequestURI()

	u, err := ctx.Request().URL.Parse(url)
	if err != nil {
		fmt.Println(err)
	}
	u.Path = ""
	u.RawQuery = ""
	u.Fragment = ""

	return u.String()
}
