package series

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"series-api/helpers"
	"series-api/models"
	"series-api/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	url := helpers.GetEnv("SERIES_API_URL", "error")
	apiKey := helpers.GetEnv("SERIES_API_KEY", "error")
	pin := helpers.GetEnv("SERIES_PIN", "error")

	if url == "error" || apiKey == "error" || pin == "error" {
		utils.RespondWithError(c, errors.New("internal server error"), http.StatusUnauthorized)
		return
	}

	body := models.LoginBodyRequest{
		ApiKey: apiKey,
		Pin:    pin,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		utils.RespondWithError(c, err, http.StatusUnauthorized)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonBody))
	if err != nil {
		utils.RespondWithError(c, err, http.StatusUnauthorized)
		return
	}
	defer resp.Body.Close()

	var response models.LoginResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		utils.RespondWithError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   response.Data,
	})

}
