package series

import (
	"encoding/json"
	"net/http"
	"series-api/helpers"
	"series-api/models"

	"github.com/gin-gonic/gin"
)

func GetSeries(c *gin.Context, page string, accessToken string) {

	url := helpers.GetEnv("SERIES_API_URL", "error")

	if url == "error" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "KO",
			"error":  "Internal server error",
		})
	}

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "KO",
			"error":  "Internal server error",
		})
	}

	defer resp.Body.Close()

	var response models.SeriesResponse

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
