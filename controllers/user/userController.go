package controllers

import (
	"SleekSpace/middleware"
	userService "SleekSpace/services/user"

	"github.com/gin-gonic/gin"
)

func GetUsers(router *gin.Engine) {
	router.GET("/users", userService.GetUsers)
}

func UpdateUser(router *gin.Engine) {
	router.PUT("/user/:id", userService.UpdateUser)
}

func GetUser(router *gin.Engine) {
	router.GET("/user/:id", middleware.AuthValidator, userService.GetUser)
}

func GetUserByEmail(router *gin.Engine) {
	router.GET("/user", middleware.AuthValidator, userService.GetUserByEmail)
}

func DeleteUser(router *gin.Engine) {
	router.DELETE("/user/:id", userService.DeleteUser)
}
