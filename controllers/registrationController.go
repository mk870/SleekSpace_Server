package controllers

import (
	"SleekSpace/services"

	"github.com/gin-gonic/gin"
)

func Registration(router *gin.Engine) {
	router.POST("/register", services.Registration)
}
