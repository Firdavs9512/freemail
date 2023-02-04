package test__test

import (
	"testing"

	"github.com/firdavs9512/freemail/services/freemail/components"
)

func TestSendEmail(t *testing.T) {
	_ = components.SendEmail("firdavs20fe@gmail.com", "ncjas", "test for test", "")
}
