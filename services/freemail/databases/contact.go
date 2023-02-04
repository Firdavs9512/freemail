package databases

// Create new contact email
func CreateContactMessage(name, email, message string) ContactEmails {
	new := ContactEmails{
		FullName: name,
		Email:    email,
		Message:  message,
	}

	res := db.Create(&new)

	if res.Error != nil {
		return ContactEmails{}
	}

	sub := Subscribles{
		Email: email,
	}
	err := db.Create(&sub)

	if err != nil {
		return ContactEmails{}
	}

	return new
}

// Delete contact email
func DeleteContactEmail(id uint) error {
	res := db.Delete(&ContactEmails{}, id)
	return res.Error
}

// Get all contact emails
func AllContactEmails() []ContactEmails {
	var emails []ContactEmails

	db.Find(&emails)

	return emails
}

// Get one email
func GetOneEmail(id uint) ContactEmails {
	var emails ContactEmails
	db.First(&emails, id)
	return emails
}
