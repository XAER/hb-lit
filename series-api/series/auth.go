package series

import (
	"bytes"
	"encoding/json"
	"net/http"
	"series-api/helpers"
	"series-api/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	url := helpers.GetEnv("SERIES_API_URL", "error")
	apiKey := helpers.GetEnv("SERIES_API_KEY", "error")

	if url == "error" || apiKey == "error" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "KO",
			"error":  "Internal server error",
		})
	}

	body := models.LoginBodyRequest{
		ApiKey: apiKey,
		Pin:    "IPQZOCMN",
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "KO",
			"error":  "Internal server error",
		})
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "KO",
			"error":  "Internal server error",
		})
	}
	defer resp.Body.Close()

	var response models.LoginResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "KO",
			"error":  "Internal server error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   response.Data,
	})

}
