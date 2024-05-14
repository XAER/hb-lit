package utils

import (
	"series-api/models"

	"github.com/gin-gonic/gin"
)

func RespondWithError(c *gin.Context, err error, statusCode int) {
	c.JSON(statusCode, models.SErrorResponse{
		Status:  "KO",
		Message: err.Error(),
	})
}
