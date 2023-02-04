package controller

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/firdavs9512/freemail/services/freemail/components"
	"github.com/firdavs9512/freemail/services/freemail/databases"
	"github.com/kataras/iris/v12"
)

func AdminCreateTemplateReq(ctx iris.Context) {

	// Create directory for upload files
	components.CreateDirectory("uploads")
	ctx.SetMaxRequestBodySize(30 * iris.MB)

	name := ctx.PostValue("name")
	primum, _ := ctx.PostValueBool("primum")
	active, _ := ctx.PostValueBool("active")

	// Upload template html file
	_, fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.SetCookieKV("error", "Error for upload file!")
		ctx.Redirect("/admin/create-template")
		return
	}
	components.CreateDirectory("uploads/shablon")
	filename := fmt.Sprintf("%s%s", components.RandomText(10), fileHeader.Filename)
	dest := filepath.Join("./uploads/shablon", filename)
	ctx.SaveFormFile(fileHeader, dest)

	// Upload template example image
	components.CreateDirectory("uploads/images")
	_, imageHeader, err := ctx.FormFile("image")
	if err != nil {
		ctx.SetCookieKV("error", "Error for upload file!")
		ctx.Redirect("/admin/create-template")
		return
	}
	imagename := fmt.Sprintf("%s%s", components.RandomText(10), imageHeader.Filename)
	desi := filepath.Join("./uploads/images", imagename)
	ctx.SaveFormFile(imageHeader, desi)

	// Create new template table for base
	tmpl, err := databases.TemplateCreate(name, filename, imagename, primum, active)

	if err != nil {
		ctx.SetCookieKV("error", "Error for create template!")
		ctx.Redirect("/admin/create-template")
		return
	}

	if tmpl.File == "" {
		ctx.SetCookieKV("error", "Error for create template!")
		ctx.Redirect("/admin/create-template")
		return
	}

	// Full complite return success message
	ctx.SetCookieKV("message", "Template successfull created!")
	ctx.Redirect("/admin")

}

// Get all templates
func GetAllTemplates() []databases.Templates {
	temp := databases.GetTemplates()
	return temp
}

// Delete template
func AdminDeleteTemplate(ctx iris.Context) {
	id, _ := ctx.PostValueUint("id")

	temp := databases.GetTemplate(id)

	if temp.Name == "" {
		ctx.SetCookieKV("error", "Error find template!")
		ctx.Redirect("/admin/templates")
	}

	// delete image and file
	imagename := fmt.Sprintf("./uploads/images/%s", temp.Image)
	errs := os.Remove(imagename)
	if errs != nil {
		ctx.SetCookieKV("error", "Error delete image!")
		ctx.Redirect("/admin/templates")
	}

	filename := fmt.Sprintf("./uploads/shablon/%s", temp.File)
	errss := os.Remove(filename)
	if errss != nil {
		ctx.SetCookieKV("error", "Error delete file!")
		ctx.Redirect("/admin/templates")
	}

	err := databases.TemplateDelete(id)
	if err != nil {
		ctx.SetCookieKV("error", "Error delete template!")
		ctx.Redirect("/admin/templates")
	}

	ctx.SetCookieKV("message", "Template success deleted!")
}

// Template update active or primum
func AdminTemplateUpdate(ctx iris.Context) {
	id, _ := ctx.PostValueUint("id")
	active, _ := ctx.PostValueBool("active")
	primum, _ := ctx.PostValueBool("primum")

	temp := databases.GetTemplate(id)

	if temp.Name == "" {
		ctx.SetCookieKV("error", "Template update error!")
		return
	}

	err := databases.TemplateUpdate(id, temp.Name, temp.File, primum, active)

	if err != nil {
		ctx.SetCookieKV("error", "Template update error!")
		return
	}

	ctx.SetCookieKV("message", "Template success updated!")

}
