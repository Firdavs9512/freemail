package test__test

import (
	"testing"

	"github.com/firdavs9512/freemail/services/freemail/databases"
)

var name string = "cbhdshiabclusbdfisauflausg"
var pass string = "jcbsalichblsaclsybclyablcbs"

func TestAdminCreate(t *testing.T) {

	databases.BaseInit()
	result := databases.AdminCreate(name, pass, true)

	if result != nil {
		t.Errorf(result.Error())
	}

}

func TestAdminDisable(t *testing.T) {
	res := databases.AdminUpdate(1, true)
	if res != nil {
		t.Errorf(res.Error())
	}
}

func TestAdminDelete(t *testing.T) {
	res := databases.AdminDelete(1)
	if res != nil {
		t.Errorf(res.Error())
	}
}
