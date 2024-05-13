package config

import "github.com/gin-gonic/gin"

var App *Services

type Services struct {
	Router *gin.Engine
	Config *GinConfig
	MODE   string
}

func InitializeServices(router *gin.Engine, config *GinConfig) *Services {
	return &Services{
		Router: router,
		Config: config,
		MODE:   gin.Mode(),
	}
}

func (s *Services) GetMode() string {
	return s.MODE
}
