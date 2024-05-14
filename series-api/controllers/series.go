package controllers

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetSeriesController(c *gin.Context) {
	// getting the parameters
	params := c.Request.URL.Query()

	// getting the page parameter
	page := params.Get("page")

	if page == "" {
		page = "0"
	}

	accessToken := c.Request.Header.Get("Authorization")

	accessToken = strings.Split(accessToken, "Bearer ")[1]

}
