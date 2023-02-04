package main

import (
	"fmt"

	"github.com/firdavs9512/freemail/services/freemail/databases"
)

func main() {
	db := databases.Connect()
	fmt.Println("**** Databases migration runned ****")
	db.AutoMigrate(
		&databases.Users{},
		&databases.Admins{},
		&databases.Token{},
		&databases.Templates{},
		&databases.Subscribles{},
		&databases.ContactEmails{},
	)
	// db.AutoMigrate(&databases.Token{})
	databases.BaseInit()

	databases.Create()
}
