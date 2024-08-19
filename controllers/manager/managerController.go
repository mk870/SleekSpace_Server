package controllers

import (
	"SleekSpace/middleware"
	managerService "SleekSpace/services/manager"
	standService "SleekSpace/services/property/stand"

	"github.com/gin-gonic/gin"
)

func CreateManager(router *gin.Engine) {
	router.POST("/manager", middleware.AuthValidator, managerService.CreateManager)
}

func GetManagerByUserId(router *gin.Engine) {
	router.GET("/manager/user/:userId", middleware.AuthValidator, managerService.GetManagerByUserId)
}

func GetManagerById(router *gin.Engine) {
	router.GET("/manager/:id", middleware.AuthValidator, managerService.GetManagerByManagerId)
}

func GetManagerStandsByManagerId(router *gin.Engine) {
	router.GET("/manager/stands/:id", middleware.AuthValidator, standService.GetManagerStandsByManagerId)
}

func UpdateManager(router *gin.Engine) {
	router.PUT("/manager/:id", middleware.AuthValidator, managerService.UpdateManagerEmailAndName)
}

func DeleteManager(router *gin.Engine) {
	router.DELETE("/manager/:id", middleware.AuthValidator, managerService.DeleteManager)
}
