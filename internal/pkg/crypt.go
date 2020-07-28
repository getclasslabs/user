package pkg

import (
	"golang.org/x/crypto/bcrypt"
)

func Crypt(text string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

