package databases

import (
	"errors"

	"github.com/firdavs9512/freemail/services/freemail/components"
)

// New Token created function
func CreateToken(email string, name string, temp string) (Token, error) {
	var user Users

	db.First(&user, "email = ?", email).Where("deleted_at = null")

	if user.Email == "" {
		return Token{}, errors.New("User is deleted!")
	}

	var tokens []Token
	db.Find(&tokens, "users_id = ?", user.ID)

	if user.Plan == 1 && len(tokens) >= 3 {
		return Token{}, errors.New("User token limit full !!!")
	}
	if user.Plan == 2 && len(tokens) >= 5 {
		return Token{}, errors.New("User token limit max !!!")
	}
	if user.Plan == 3 && len(tokens) >= 7 {
		return Token{}, errors.New("User token limit full !!!")
	}
	if user.Plan == 4 && len(tokens) >= 10 {
		return Token{}, errors.New("User token limit full !!!")
	}

	randomurl := components.RandomText(30)

	new := Token{
		Name:     name,
		Maxreq:   10000,
		Request:  0,
		Template: temp,
		Url:      randomurl,
		UsersID:  user.ID,
	}
	result := db.Create(&new)
	return new, result.Error
}

// Token deleted function
func DeleteToken(tokenid uint) error {
	res := db.Delete(&Token{}, tokenid)
	return res.Error
}

// Token template update function
func UpdateTokenTemplate(tokenid, templateid uint) error {
	res := db.Model(&Token{}).Where("id = ?", tokenid).Update("template", templateid)
	return res.Error
}

// Token request update funtion
func TokenRequestUpdate(tokenid, req uint) error {
	res := db.Model(&Token{}).Where("id = ?", tokenid).Update("request", req)
	return res.Error
}

// Token full delete function
func TokenFullDelete(tokenid, userid uint) error {

	ress := db.Unscoped().Delete(&Token{}, "id = ? and users_id = ?", tokenid, userid)
	return ress.Error
}

// Get Full user tokens
func GetFullTokens(userid uint) []Token {
	var user Users
	db.First(&user, userid)
	var tokens []Token
	db.Where("users_id = ?", user.ID).Find(&tokens)

	// db.Find(&tokens, "users_id = ?", userid)
	return tokens
}

// Get data for token url
func GetDataForUrl(url string) Token {
	var token Token

	db.First(&token, "url = ?", url)

	return token
}

// Get all tokens
func GetAllTokens() []Token {
	var token []Token

	db.Find(&token)
	return token
}
