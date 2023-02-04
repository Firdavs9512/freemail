package main

import (
	"bytes"
	"html/template"

	"gopkg.in/gomail.v2"
)

func main() {
	// Load the HTML template into a string
	tmpl, err := template.ParseFiles("./template.html")
	if err != nil {
		panic(err)
	}

	// Create a new message
	m := gomail.NewMessage()
	m.SetHeader("From", "firdavs@example.com")
	m.SetHeader("To", "firdavsharipovredmi8@gmail.com")
	m.SetHeader("Subject", "Subject of the email")

	// Render the HTML template and set it as the body of the message
	var buf bytes.Buffer
	data := map[string]interface{}{
		"Message": "This is the message you want to display",
	}
	if err := tmpl.Execute(&buf, data); err != nil {
		panic(err)
	}
	m.SetBody("text/html", buf.String())

	// Connect to the SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, "firdavsreadmi8@gmail.com", "utmayhmeahyvrehs")

	// Send the message
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
