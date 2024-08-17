package controllers

import (
	"SleekSpace/middleware"
	managerService "SleekSpace/services/manager"

	"github.com/gin-gonic/gin"
)

func UpdateManagerContactNumbers(router *gin.Engine) {
	router.PUT("/manager/contacts/:id", middleware.AuthValidator, managerService.UpdateManagerContactNumbers)
}
