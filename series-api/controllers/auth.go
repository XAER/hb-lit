package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"series-api/logger"
	"series-api/models"
	"series-api/series"
	"series-api/utils"

	"github.com/gin-gonic/gin"
)

func LoginController(c *gin.Context) {
	var loginRequest models.LoginBodyStruct

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		utils.RespondWithError(c, err, http.StatusInternalServerError)
		return
	}

	userEmail := loginRequest.Email
	userPassword := loginRequest.Password

	if userEmail == "" || userPassword == "" {
		utils.RespondWithError(c, errors.New("email and password are required"), http.StatusUnauthorized)
		return
	}

	logger.Info(fmt.Sprintf("user email from request: %s", userEmail))

	isPasswordCorrect, err := utils.CheckPassword(userPassword)
	if err != nil || !isPasswordCorrect {
		utils.RespondWithError(c, errors.New("the credentials you entered did not match any registered profile"), http.StatusNotFound)
		return
	}

	// If password is correct, I have to send a POST request to the Movies/Series API to retrive the bearer token
	// and then send it back to the user

	if isPasswordCorrect {
		series.Login(c)
	}
}
