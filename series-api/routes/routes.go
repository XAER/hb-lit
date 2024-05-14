package routes

import (
	"series-api/controllers"
	"series-api/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 15 * time.Second,
	}))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", controllers.PingController)

		auth := v1.Group("/auth")
		{
			auth.POST("/login", controllers.LoginController)
		}
		tv := v1.Group("/tv")
		{
			tv.GET("/series", middleware.TokenAuthMiddleware(), controllers.GetSeriesController)
		}
	}

	return r
}
