package utils

import (

	"golang.org/x/crypto/bcrypt"
)

func BcryptHash(passKey string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passKey), 14)
	return string(bytes), err
}