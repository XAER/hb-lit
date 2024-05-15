package main

import (
	"fmt"
	"log"
	"series-api/config"
	"series-api/helpers"
	"series-api/logger"
	"series-api/routes"

	"github.com/gin-gonic/gin"
)

var (
	ginConfig config.GinConfig
	router    *gin.Engine
)

func initGinServer() {
	ginConfig.SERVER_NAME = helpers.GetEnv("SERVER_NAME", "ERP_BACKEND")
	ginConfig.SERVER_PORT = helpers.GetEnv("SERVER_PORT", "9090")
	ginConfig.SERVER_ENV = helpers.GetEnv("SERVER_ENV", "dev")

	config.SetGinMode(ginConfig.SERVER_ENV)
	router = routes.SetupRouter()
}

func main() {
	initGinServer()
	logger.InitializeLogger()

	config.App = config.InitializeServices(router, &ginConfig)

	log.Fatal(router.Run(fmt.Sprintf("0.0.0.0:%s", ginConfig.SERVER_PORT)))
}
