package routers

import (
	"github.com/OtavioValadao/api-main-go.git/config"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	// Run the server
	router.Run(config.GetEnvironments("API_PORT"))
}
