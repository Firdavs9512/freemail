package controller

import "github.com/kataras/iris/v12"

// Home index page controller
func Index(ctx iris.Context) {
	ctx.View("home/index.html")
}

// Home pricing page controller
func Pricing(ctx iris.Context) {
	ctx.View("home/pricing-plan.html")
}

// Home privacy page controller
func Privacy(ctx iris.Context) {
	ctx.View("home/privacy.html")
}

// Home about page controller
func About(ctx iris.Context) {
	ctx.View("home/about.html")
}

// Home contact page controller
func Contact(ctx iris.Context) {
	ctx.ViewData("message", ctx.GetCookie("message"))
	ctx.RemoveCookie("message")
	ctx.View("home/contact.html")
}

// Home register page controller
func ResetPassword(ctx iris.Context) {
	ctx.View("dashboard/auth-reset-password.html")
}

// Home register page controller
func ForgotPassword(ctx iris.Context) {
	ctx.View("dashboard/auth-forgot-password.html")
}

// Home service page controller
func Services(ctx iris.Context) {
	ctx.View("home/service.html")
}

// Home ServicesPromotion page controller
func ServicesPromotion(ctx iris.Context) {
	ctx.View("home/service-promotion.html")
}

// Home Reseller page controller
func Reseller(ctx iris.Context) {
	ctx.View("home/reseller.html")
}


// robots.tx
func Robotstxt(ctx iris.Context) {
	ctx.ServeFile("template/home/robots.txt")
}

// sitemap.xml
func Sitemapxml(ctx iris.Context) {
	ctx.ServeFile("template/home/sitemap.xml")
}

// Google ads
func Adstxt(ctx iris.Context){
	ctx.ServeFile("template/home/ads.txt")
}
