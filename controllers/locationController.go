package controllers

import (
	"SleekSpace/middleware"
	"SleekSpace/services"

	"github.com/gin-gonic/gin"
)

func CreateLocation(router *gin.Engine) {
	router.POST("/location", middleware.AuthValidator, services.CreateLocation)
}

func UpdateLocation(router *gin.Engine) {
	router.PUT("/location", middleware.AuthValidator, services.UpdateLocation)
}
