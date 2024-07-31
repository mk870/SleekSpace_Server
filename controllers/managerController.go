package controllers

import (
	"SleekSpace/middleware"
	"SleekSpace/services"

	"github.com/gin-gonic/gin"
)

func CreateManager(router *gin.Engine) {
	router.POST("/manager", middleware.AuthValidator, services.CreateManager)
}

func GetManagerByUserId(router *gin.Engine) {
	router.GET("/manager/user/:userId", middleware.AuthValidator, services.GetManagerByUserId)
}

func GetManagerById(router *gin.Engine) {
	router.GET("/manager/:id", middleware.AuthValidator, services.GetManagerByManagerId)
}

func UpdateManager(router *gin.Engine) {
	router.PUT("/manager/:id", middleware.AuthValidator, services.UpdateManager)
}

func DeleteManager(router *gin.Engine) {
	router.DELETE("/manager/:id", middleware.AuthValidator, services.DeleteManager)
}
