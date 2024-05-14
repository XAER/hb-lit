package utils

import (
	"errors"
	"series-api/helpers"
)

func CheckPassword(password string) (bool, error) {
	// Getting the password from the .env file
	userPassword := helpers.GetEnv("TEMP_PASSWORD", "password")

	if userPassword == "password" {
		return false, errors.New("wrong password")
	}
	if password == userPassword {
		return true, nil
	} else {
		return false, errors.New("wrong password")
	}
}
