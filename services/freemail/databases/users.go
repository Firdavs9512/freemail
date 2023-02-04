package databases

import (
	"time"
)

// User create function
func UserCreate(email, pass, last, first string) (Users, error) {
	user := Users{
		LastName:  last,
		FirstName: first,
		Email:     email,
		Password:  pass,
		Plan:      1,
		DatePlan:  time.Now(),
	}

	result := db.Create(&user)
	return user, result.Error
}

// User login function
func UserLogin(email string) Users {
	var user Users

	db.First(&user, "email = ?", email)

	return user

}

// User delete function
func UserDelete(id uint) error {
	result := db.Delete(&Users{}, id)
	return result.Error
}

// User plan update function
func UserPlanUpdate(id uint, plan int) error {
	var user Users
	db.First(&user, id)
	user.DatePlan = time.Now()
	user.Plan = plan
	result := db.Save(&user)

	return result.Error
}

// Full user delete function
func FullUserDelete(userid uint) error {
	res := db.Unscoped().Delete(&Users{}, userid)
	return res.Error
}

// user email controll
func UserControll(email string) bool {
	res := db.First(&Users{}, "email = ?", email)
	if res.Error != nil {
		return true
	}
	return false
}

// Get all users
func GetAllUsers() []Users {
	var users []Users

	db.Find(&users)

	return users

}

// Get All Primum users
func GetPrimumUsers() []Users {
	var users []Users

	db.Find(&users, "plan!=1")
	return users
}
