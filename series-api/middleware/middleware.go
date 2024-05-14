package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Add(r *gin.Engine, h ...gin.HandlerFunc) {
	for _, v := range h {
		r.Use(v)
	}
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "KO",
				"error":  "Unauthorized: Token is missing",
			})
			c.Abort()
			return
		}

		splitToken := strings.Split(token, "Bearer ")
		if len(splitToken) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "KO",
				"error":  "Unauthorized: Token is invalid",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
