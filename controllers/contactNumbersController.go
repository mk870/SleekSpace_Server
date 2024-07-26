package controllers

import (
	"SleekSpace/middleware"
	"SleekSpace/services"

	"github.com/gin-gonic/gin"
)

func CreateContact(router *gin.Engine) {
	router.POST("/contact-number", middleware.AuthValidator, services.CreateContactNumber)
}

func UpdateContactNumbers(router *gin.Engine) {
	router.PUT("/contact-number/:id", middleware.AuthValidator, services.UpdateContactNumbers)
}
