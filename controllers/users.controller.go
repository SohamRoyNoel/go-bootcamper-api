package controllers;

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HealthCheck() gin.HandlerFunc {
	return func (c *gin.Context)  {
		fmt.Println("reached controller")
	}
}
