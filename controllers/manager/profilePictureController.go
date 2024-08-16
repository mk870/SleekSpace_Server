package controllers

import (
	"SleekSpace/middleware"
	managerService "SleekSpace/services/manager"

	"github.com/gin-gonic/gin"
)

func UpdateManagerProfilePicture(router *gin.Engine) {
	router.PUT("/manager/profile-picture/:id", middleware.AuthValidator, managerService.UpdateManagerProfilePicture)
}
