package components

import (
	"log"
	"os"
)

func CreateDirectory(name string) {
	if _, err := os.Stat(name); os.IsNotExist(err) {
		err = os.Mkdir(name, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}
