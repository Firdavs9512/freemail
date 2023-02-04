package test__test

import (
	"testing"

	"github.com/firdavs9512/freemail/services/freemail/databases"
)

func TestCreateTemplate(t *testing.T) {
	databases.BaseInit()

	temp, err := databases.TemplateCreate("test", "test", "test", true, true)

	if err != nil {
		t.Errorf("Error for create new template! error: %s", err.Error())
	}

	if temp.File == "" {
		t.Errorf("Error create new template! Template is empty!")
	}

	if temp.Name == "" {
		t.Errorf("Error create new template! Template is empty!")
	}
}

func TestUpdateTemplate(t *testing.T) {
	databases.BaseInit()

	res := databases.TemplateUpdate(1, "test1", "file", true, false)

	if res != nil {
		t.Errorf("Error for template update function! error: %s", res.Error())
	}
}
func TestTemplatedelete(t *testing.T) {
	databases.BaseInit()

	res := databases.TemplateDelete(1)

	if res != nil {
		t.Errorf("Error for delete template function! error %s", res.Error())
	}
}
