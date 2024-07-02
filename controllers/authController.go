package controllers

import (
	"SleekSpace/middleware"
	"SleekSpace/services"

	"github.com/gin-gonic/gin"
)

func Login(router *gin.Engine) {
	router.POST("/login", services.Login)
}

func LoginOut(router *gin.Engine) {
	router.GET("/logout", middleware.AuthValidator, services.Logout)
}
