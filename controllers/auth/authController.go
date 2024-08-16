package controllers

import (
	"SleekSpace/middleware"
	userService "SleekSpace/services/user"

	"github.com/gin-gonic/gin"
)

func Login(router *gin.Engine) {
	router.POST("/login", userService.Login)
}

func LoginOut(router *gin.Engine) {
	router.GET("/logout", middleware.AuthValidator, userService.Logout)
}
