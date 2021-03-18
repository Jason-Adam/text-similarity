package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jason-adam/text-similarity/internal/api/config"
	"github.com/jason-adam/text-similarity/internal/api/router"
)

func setConfiguration(configPath string) {
	config.Setup(configPath)
	gin.SetMode(config.GetConfig().Server.Mode)
}

// Run intializes and runs the application
func Run(configPath string) {
	if configPath == "" {
		configPath = "data/config.yml"
	}
	setConfiguration(configPath)
	conf := config.GetConfig()
	web := router.Setup()
	log.Println("text similarity running on port " + conf.Server.Port)
	_ = web.Run(":" + conf.Server.Port)
}
