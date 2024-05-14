package series

import (
	"encoding/json"
	"errors"
	"net/http"
	"series-api/helpers"
	"series-api/models"
	"series-api/utils"

	"github.com/gin-gonic/gin"
)

func GetSeries(c *gin.Context, page string, accessToken string) {

	url := helpers.GetEnv("SERIES_API_URL", "error")

	if url == "error" {
		utils.RespondWithError(c, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}

	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		utils.RespondWithError(c, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}

	var response models.SeriesResponse

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		utils.RespondWithError(c, errors.New("internal server error"), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"data":   response.Data,
	})

}
