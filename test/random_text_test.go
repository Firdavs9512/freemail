package test__test

import (
	"testing"

	"github.com/firdavs9512/freemail/services/freemail/components"
)

func TestRandomtext(t *testing.T) {
	new := components.RandomText(10)

	if len(new) == 0 && new == "" {
		t.Errorf("Error for component random text function")
	}
}
