package middleware

import (
	"errors"
	"net/http"
	"series-api/utils"
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
			utils.RespondWithError(c, errors.New("unauthorized: Token is missing"), http.StatusUnauthorized)
			c.Abort()
			return
		}

		splitToken := strings.Split(token, "Bearer ")
		if len(splitToken) != 2 {
			utils.RespondWithError(c, errors.New("unauthorized: Token is invalid"), http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Next()
	}
}
