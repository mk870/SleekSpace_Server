package controllers

import (
	userService "SleekSpace/services/user"

	"github.com/gin-gonic/gin"
)

func PasswordUpdate(router *gin.Engine) {
	router.PUT("/password", userService.UpdatePassword)
}
