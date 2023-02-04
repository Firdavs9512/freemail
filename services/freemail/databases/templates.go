package databases

// Template create function
func TemplateCreate(name, file, image string, primum, active bool) (Templates, error) {
	new := Templates{
		Name:     name,
		File:     file,
		Image:    image,
		Active:   active,
		IsPrimum: primum,
	}

	res := db.Create(&new)
	return new, res.Error
}

// Template rename function
func TemplateUpdate(id uint, name, file string, primum, active bool) error {
	var template Templates
	db.First(&template, id)
	template.Active = active
	template.File = file
	template.IsPrimum = primum
	template.Name = name

	res := db.Save(&template)

	return res.Error
}

// Template delete function
func TemplateDelete(id uint) error {
	res := db.Unscoped().Delete(&Templates{}, id)

	return res.Error
}

// Get one template
func GetTemplate(id uint) Templates {
	var temp Templates

	db.First(&temp, id)
	return temp
}

// Get all templates
func GetTemplates() []Templates {
	var tmpl []Templates
	db.Find(&tmpl, "active=true")
	return tmpl
}
