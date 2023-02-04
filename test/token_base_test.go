package test__test

import (
	"testing"

	"github.com/firdavs9512/freemail/services/freemail/databases"
)

func TestTokenCreate(t *testing.T) {
	databases.BaseInit()
	email := "test@test.com"
	check := "testfile.html"

	token, err := databases.CreateToken(email, "testtoken", check)

	if err != nil {
		t.Errorf("Error for create token function !!! error: %s", err.Error())
	}

	if token.Url == "" {
		t.Error("Error for create token function !!!")
	}
}

func TestTokenDelete(t *testing.T) {
	databases.BaseInit()

	err := databases.DeleteToken(1)

	if err != nil {
		t.Errorf("Error for delete token! error: %s", err.Error())
	}

}

func TestTokenFullDelete(t *testing.T) {
	databases.BaseInit()

	err := databases.TokenFullDelete(1, 1)

	if err != nil {
		t.Errorf("Error for delete token! error: %s", err.Error())
	}
}

func TestUpdateTokenTemplate(t *testing.T) {
	databases.BaseInit()

	res := databases.UpdateTokenTemplate(3, 3)

	if res != nil {
		t.Errorf("Error for Token template update! error: %s", res.Error())
	}
}
