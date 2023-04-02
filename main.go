package main

import (
	"github.com/MuShaf-NMS/sigmatech-test/config"
	"github.com/MuShaf-NMS/sigmatech-test/database"
	"github.com/MuShaf-NMS/sigmatech-test/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// config.LoadEnv()
	config := config.LoadConfig()
	db := database.CreateConnection(config)
	defer database.CloseConnection(db)
	if config.App_Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	server := gin.Default()
	router.IntializeRouter(server, db, config)
	server.Run(":" + config.App_Port)
}
