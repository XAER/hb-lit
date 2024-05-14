package controllers

import "github.com/gin-gonic/gin"

func GetSeriesController(c *gin.Context) {
	// getting the parameters
	params := c.Request.URL.Query()

	// getting the page parameter
	page := params.Get("page")

	if page == "" {
		page = "0"
	}

	c.JSON(200, gin.H{
		"message": "Get series",
	})
}
