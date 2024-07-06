package controllers

import (
	"SleekSpace/services"

	"github.com/gin-gonic/gin"
)

func PasswordUpdate(router *gin.Engine) {
	router.PUT("/password", services.UpdatePassword)
}
