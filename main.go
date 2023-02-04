package main

import (
	"log"

	"github.com/firdavs9512/freemail/routers"
	"github.com/joho/godotenv"

	"fmt"
	"os"

	"github.com/firdavs9512/freemail/services/freemail/components"
	"github.com/firdavs9512/freemail/services/freemail/databases"

	"github.com/kataras/iris/v12"
)

func init() {
	// initilizatsion for databases
	databases.BaseInit()

	// Migrate database
	// databases.Migrate()

	// loading env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {

	app := iris.Default()

	// Routers
	routers.IndexRouter(app)

	// Dashboard routers
	dashboard := app.Party("/dashboard")
	routers.DashboardRouter(dashboard)

	// Admin routers
	admin := app.Party("/admin")
	routers.AdminRouter(admin)

	// Documentatsion routers
	docs := app.Party("/docs")
	routers.DocsRouter(docs)

	// Api routers
	apiv1 := app.Party("/api/v1")
	routers.FreeMailApiV1(apiv1)

	// Html files registration
	tmpl := iris.HTML("./template", ".html")
	app.RegisterView(tmpl)
	components.Mathematik(*tmpl)

	// Static files
	app.HandleDir("/template", iris.Dir("./template/"))
	app.HandleDir("/uploads", iris.Dir("./uploads/images/"))

	// 404 page not found page
	app.OnErrorCode(iris.StatusNotFound, routers.NotFoundHandler)

	// Run application
	ports := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	app.Listen(ports)
}
