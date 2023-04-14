package routes

import (
	"bootcamp.com/server/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouters(inComingRoutes *gin.Engine) {
	inComingRoutes.GET("/api/health", controllers.HealthCheck());
}