package test__test

import (
	"testing"

	"github.com/firdavs9512/freemail/services/freemail/components"
)

func TestHashpas(t *testing.T) {
	res, err := components.HashPass("test")
	if err != nil {
		t.Errorf("Error for hash function! error: %s", err.Error())
	}

	ress, err := components.HashPass("test")
	if err != nil {
		t.Errorf("Error for hash function! error: %s", err.Error())
	}

	if res != ress {
		t.Error(res, "\n", ress)
	}
}
