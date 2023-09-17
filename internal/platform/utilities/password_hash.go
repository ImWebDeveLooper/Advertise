package utilities

import "golang.org/x/crypto/bcrypt"

func PasswordHash(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err.Error()
	}
	return string(hashedPassword)
}
