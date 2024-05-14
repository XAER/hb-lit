package controllers

import (
	"net/http"
	"series-api/models"
	"series-api/series"
	"series-api/utils"

	"github.com/gin-gonic/gin"
)

func LoginController(c *gin.Context) {
	var loginRequestBody models.LoginBodyStruct

	if err := c.ShouldBindJSON(&loginRequestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "KO",
			"message": err.Error(),
		})
	}

	userEmail := loginRequestBody.Email
	userPassword := loginRequestBody.Password

	if userEmail == "" || userPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "KO",
			"message": "Email and password are required",
		})
	}

	if userEmail == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "KO",
			"message": "Email is required",
		})
	}

	if userPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "KO",
			"message": "Password is required",
		})
	}

	isPasswordCorrect, err := utils.CheckPassword(userPassword)
	if err != nil || !isPasswordCorrect {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "KO",
			"message": err.Error(),
		})
	}

	// If password is correct, I have to send a POST request to the Movies/Series API to retrive the bearer token
	// and then send it back to the user

	if isPasswordCorrect {
		series.Login(c)
	}
}
