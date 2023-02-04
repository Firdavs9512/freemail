package test__test

import (
	"testing"

	"github.com/firdavs9512/freemail/services/freemail/databases"
)

func TestMysqlConnect(t *testing.T) {
	conn := databases.Connect()

	if conn.Error != nil {
		t.Errorf("Error for connection for database !!!")
	}
}
