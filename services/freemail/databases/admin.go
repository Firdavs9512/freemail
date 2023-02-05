package databases

import (
	"fmt"
)

// Admin create function
func AdminCreate(name string, pass string, acc bool) error {
	db := Connect()
	new := Admins{
		Username: name,
		Password: pass,
		Active:   acc,
	}

	result := db.Create(&new)
	return result.Error
}

// Admin Update function
func AdminUpdate(id uint, active bool) error {
	result := db.Model(&Admins{}).Where("id = ?", id).Update("active", active)
	return result.Error
}

// Admin deleted function
func AdminDelete(id uint) error {
	result := db.Delete(&Admins{}, "id = ?", id)
	return result.Error
}

// Auto created admin
func Create() {
	admin := Admins{
		Username: "admin",
		Password: "test",
		Active:   true,
		Secret:   true,
	}

	result := db.Create(&admin)
	fmt.Println(result.Error)
}

func AdminLogin(name, pass string) Admins {
	var admin Admins
	db.Where("active=true AND username = ? AND password = ?", name, pass).Find(&admin)

	return admin

}

func AdminLoginForName(name string) Admins {
	var admin Admins
	db.First(&admin, "username = ? and active=true", name)

	return admin

}

// Get all admins
func GetAllAdmins() []Admins {
	var admins []Admins

	db.Find(&admins, "secret=false")

	return admins

}
