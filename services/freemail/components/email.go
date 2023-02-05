package components

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"gopkg.in/gomail.v2"
)

func SendEmail(to_email, message, templ, subject string) error {

	email := os.Getenv("EMAIL_ADDRES")
	password := os.Getenv("EMAIL_PASSWORD")

	// Load the HTML template into a string
	tempfile := fmt.Sprintf("uploads/shablon/%s", templ)
	tmpl, err := template.ParseFiles(tempfile)
	if err != nil {
		return err
	}

	// Create a new message
	m := gomail.NewMessage()
	m.SetHeader("From", email)
	m.SetHeader("To", to_email)
	m.SetHeader("Subject", subject)

	// Render the HTML template and set it as the body of the message
	var buf bytes.Buffer
	data := map[string]interface{}{
		"message": message,
	}
	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}
	m.SetBody("text/html", buf.String())

	// Connect to the SMTP server
	d := gomail.NewDialer("smtp.gmail.com", 587, email, password)

	// Send the message
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
