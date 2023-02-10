package utils

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashForPassword(input string) (string, error) {
	password := []byte(input)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(input string, password string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(input), []byte(password))
	return err == nil, err
}

func GenerateUUID() string {
	return uuid.New().String()

}
