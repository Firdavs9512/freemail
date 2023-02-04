package test__test

import (
	"testing"

	"github.com/firdavs9512/freemail/services/freemail/databases"
)

var user_id uint

func TestUserCreate(t *testing.T) {
	databases.BaseInit()
	lastname := "Test"
	firstname := "testt"
	email := "test@testing.example"
	pass := "bsachsbclhbldhvbl"
	user, err := databases.UserCreate(email, pass, lastname, firstname)

	if err != nil {
		t.Error("Error for create new user !!!")
	}

	if user.ID == 0 {
		t.Error("Error for create new user !!!")
	}

	if user.Email != email {
		t.Error("Error for create new user !!!")
	}

	user_id = user.ID
}

func TestUserUpdatePlan(t *testing.T) {
	var plan int = 3
	databases.BaseInit()

	err := databases.UserPlanUpdate(user_id, plan)

	if err != nil {
		t.Error("Error for User updata plan function !!!")
	}

}

func TestUserDelete(t *testing.T) {
	databases.BaseInit()

	err := databases.UserDelete(user_id)

	if err != nil {
		t.Error("Error for user delete function !!!")
	}
}
func TestFullUserDelete(t *testing.T) {
	databases.BaseInit()

	err := databases.FullUserDelete(1)

	if err != nil {
		t.Errorf("Error for full delete users function! error: %s", err.Error())
	}
}
