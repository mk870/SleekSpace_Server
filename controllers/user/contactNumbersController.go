package controllers

import (
	"SleekSpace/middleware"
	userService "SleekSpace/services/user"

	"github.com/gin-gonic/gin"
)

func CreateContact(router *gin.Engine) {
	router.POST("/contact-number", middleware.AuthValidator, userService.CreateContactNumber)
}

func UpdateContactNumbers(router *gin.Engine) {
	router.PUT("/contact-number/:id", middleware.AuthValidator, userService.UpdateContactNumbers)
}
