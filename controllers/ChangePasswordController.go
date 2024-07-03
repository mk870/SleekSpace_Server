package controllers

import (
	"SleekSpace/services"

	"github.com/gin-gonic/gin"
)

func ChangePassword(router *gin.Engine) {
	router.POST("/password", services.ChangePassword)
}

func PasswordChangeAndVerification(router *gin.Engine) {
	router.PUT("/password", services.UpdatePasswordAndCodeVerification)
}
