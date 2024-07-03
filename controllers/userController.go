package controllers

import (
	"SleekSpace/middleware"
	"SleekSpace/services"

	"github.com/gin-gonic/gin"
)

func GetUsers(router *gin.Engine) {
	router.GET("/users", services.GetUsers)
}

func UpdateUser(router *gin.Engine) {
	router.PUT("/user/:id", services.UpdateUser)
}

func GetUser(router *gin.Engine) {
	router.GET("/user/:id", middleware.AuthValidator, services.GetUser)
}

func DeleteUser(router *gin.Engine) {
	router.DELETE("/user/:id", services.DeleteUser)
}
