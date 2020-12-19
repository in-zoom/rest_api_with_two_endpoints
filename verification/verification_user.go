package verification

import (
	"errors"
	"os"
	"strings"
)

const (
	messageErr string = "Введен неверный логин или пароль"
)

func VerificationUser(enteredUserName, password string) (string, error) {

	if strings.TrimSpace(enteredUserName) == "" || strings.TrimSpace(password) == "" {
		return "", errors.New(messageErr)
	}

	userName := os.Getenv("NAME_USER")

	if userName != enteredUserName {
		return "", errors.New(messageErr)

	}
	if HashPasswordUser(os.Getenv("USER_PASS")) != HashPasswordUser(password) {
		return "", errors.New(messageErr)
	}

	return userName, nil
}
