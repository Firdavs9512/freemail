package components

import "golang.org/x/crypto/bcrypt"

// Password to generate hash
func HashPass(a string) (string, error) {
	t, err := bcrypt.GenerateFromPassword([]byte(a), bcrypt.DefaultCost)
	return string(t), err
}

// Controll hash and password equel
func HashControll(a, b string) error {
	err := bcrypt.CompareHashAndPassword([]byte(a), []byte(b))

	return err
}
