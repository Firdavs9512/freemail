package databases

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Plan      int
	Tokens    []Token
	DatePlan  time.Time
}

type Admins struct {
	gorm.Model
	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Secret   bool   `gorm:"default:false"`
	Active   bool
}

type Token struct {
	gorm.Model
	Maxreq   uint
	Name     string `gorm:"not null"`
	Request  uint
	Template string `gorm:"not null"`
	Url      string `gorm:"not null;unique"`
	UsersID  uint
}

type Templates struct {
	gorm.Model
	Name     string `gorm:"not null;unique"`
	Image    string `gorm:"not null"`
	File     string `gorm:"not null;unique"`
	IsPrimum bool
	Active   bool
}

// type Paginator struct {
// 	Page      uint
// 	Total     int
// 	TotalPage int
// }

type Subscribles struct {
	gorm.Model
	Email string `gorm:"not null;unique"`
}

type ContactEmails struct {
	gorm.Model
	Email    string `gorm:"not null"`
	FullName string `gorm:"not null"`
	Message  string `gorm:"not null;type:text"`
}
