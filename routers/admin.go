package routers

import (
	"github.com/firdavs9512/freemail/services/freemail/controller"
	"github.com/kataras/iris/v12/core/router"
)

func AdminRouter(app router.Party) {
	{
		app.Get("/login", controller.AdminLogin).Name = "adminlogin"
		app.Post("/login", controller.AdminLoginRequest)
		app.Get("/logout", controller.AdminLogout).Name = "adminlogout"
		app.Get("/", controller.AdminMiddlware, controller.AdminDashboard).Name = "admin"
		app.Get("/payments", controller.AdminMiddlware, controller.AdminPayments).Name = "adminpayments"
		app.Get("/create-user", controller.AdminMiddlware, controller.AdminCreateUser).Name = "admincreateuser"
		app.Get("/users", controller.AdminMiddlware, controller.AdminAllUsers).Name = "adminusers"
		app.Get("/users/{page:uint}", controller.AdminMiddlware, controller.AdminGetUsersPaginator)
		app.Get("/create-template", controller.AdminMiddlware, controller.AdminCreateTemplate).Name = "admincreatetemplate"
		app.Get("/create-admin", controller.AdminMiddlware, controller.AdminCreateAdmin).Name = "admincreateadmin"
		app.Get("/templates", controller.AdminMiddlware, controller.AdminAllTemplates).Name = "admintemplates"
		app.Get("/admins", controller.AdminMiddlware, controller.AdminAllAdmins).Name = "admins"
		app.Get("/emails", controller.AdminMiddlware, controller.AdminContactEmails).Name = "adminemails"
		app.Get("/emails/{id:uint}", controller.AdminMiddlware, controller.AdminGetOneEmail)
		app.Get("/send-email", controller.AdminMiddlware, controller.AdminSendEmail).Name = "adminsendemail"
		app.Post("/create-template", controller.AdminMiddlware, controller.AdminCreateTemplateReq).Name = "admincreatetemplatereq"
		app.Post("/delete-template", controller.AdminMiddlware, controller.AdminDeleteTemplate).Name = "admindeletetemplate"
		app.Post("/template-update", controller.AdminMiddlware, controller.AdminTemplateUpdate).Name = "adminupdatetemplate"
		app.Post("/create-user", controller.AdminMiddlware, controller.AdminCreateUserReq).Name = "admincreateuserreq"
		app.Post("/create-admin", controller.AdminMiddlware, controller.AdminCreateAdminReq).Name = "admincreateadminreq"
		app.Post("/admin-update", controller.AdminMiddlware, controller.AdminUpdateStatus).Name = "adminupdatestatus"
		app.Post("/admin-delete", controller.AdminMiddlware, controller.AdminDeleteAdmins).Name = "admindeleteadmin"
		app.Post("/user-update", controller.AdminMiddlware, controller.AdminUserUpdate).Name = "adminuserupdate"
		app.Post("/user-delete", controller.AdminMiddlware, controller.AdminDeleteUser).Name = "admindeleteuser"
		app.Post("/contact-email", controller.ContactEmail).Name = "adminemail"
		app.Post("/contact-email-delete", controller.AdminMiddlware, controller.AdminDeleteEmail).Name = "admindeleteemail"
	}
}
