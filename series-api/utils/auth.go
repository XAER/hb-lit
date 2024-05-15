package utils

import (
	"errors"
	"series-api/helpers"
)

func CheckPassword(password string) (bool, error) {
	// Getting the password from the .env file
	userPassword := helpers.GetEnv("TEMP_PASSWORD", "error")

	if userPassword == "error" {
		return false, errors.New("wrong password")
	}

	// logger.Printf("Password from .env file: %s\n", userPassword)
	// logger.Printf("Password from request: %s\n", password)

	if password == userPassword {
		return true, nil
	} else {
		return false, errors.New("wrong password")
	}
}
