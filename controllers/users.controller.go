package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"bootcamp.com/server/configs"
	"bootcamp.com/server/handlers"
	"bootcamp.com/server/helpers"
)

func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("reached controller")
		respObj := &handlers.GeneralResponseBody{Status: 1, Message: configs.HEALTH_API_RESP}
		c.JSON(http.StatusOK, respObj)
	}
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		userHelper := helpers.RegistrationHelper("", "", ctx, cancel)
		fmt.Println("Data From Helper >>>>> ", userHelper)
	}
}
