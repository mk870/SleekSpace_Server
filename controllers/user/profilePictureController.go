package controllers

import (
	"SleekSpace/middleware"
	userService "SleekSpace/services/user"

	"github.com/gin-gonic/gin"
)

func CreateUserProfilePicture(router *gin.Engine) {
	router.POST("/user/profile-picture/:id", middleware.AuthValidator, userService.CreateUserProfilePicture)
}

func UpdateUserProfilePicture(router *gin.Engine) {
	router.PUT("/user/profile-picture/:id", middleware.AuthValidator, userService.UpdateUserProfilePicture)
}
