package controllers

import (
	"SleekSpace/middleware"
	userService "SleekSpace/services/user"

	"github.com/gin-gonic/gin"
)

func CreateLocation(router *gin.Engine) {
	router.POST("/location", middleware.AuthValidator, userService.CreateLocation)
}

func UpdateLocation(router *gin.Engine) {
	router.PUT("/location", middleware.AuthValidator, userService.UpdateLocation)
}
