package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"bootcamp.com/server/configs"
	"bootcamp.com/server/handlers"
)

func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("reached controller")
		respObj := &handlers.GeneralResponseBody{Status: 1, Message: configs.HEALTH_API_RESP}
		c.JSON(http.StatusOK, respObj)
	}
}
