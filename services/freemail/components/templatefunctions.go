package components

import (
	"github.com/kataras/iris/v12/view"
)

func Mathematik(tmpl view.HTMLEngine) {
	tmpl.AddFunc("progres", func(a, b uint) float64 {
		return float64(a) * float64(b) / 1000000
	})
	tmpl.AddFunc("residual", func(j, a int) int {
		return j - a
	})
}
