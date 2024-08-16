package controllers

import (
	userService "SleekSpace/services/user"

	"github.com/gin-gonic/gin"
)

func Registration(router *gin.Engine) {
	router.POST("/register", userService.Registration)
}
