package controllers

import (
	"SleekSpace/services"

	"github.com/gin-gonic/gin"
)

func GetAllUsersLocations(router *gin.Engine) {
	router.GET("/admin/users/location", services.GetAllUsersLocations)
}

func GetLocationById(router *gin.Engine) {
	router.GET("/admin/location/:id", services.GetLocationById)
}

func GetVerificationCodeById(router *gin.Engine) {
	router.GET("/admin/verification-code/:id", services.GetVerificationCodeById)
}

func GetAllVerificationCodes(router *gin.Engine) {
	router.GET("/admin/verification-codes", services.GetAllVerificationCodes)
}

func DeleteVerificationCode(router *gin.Engine) {
	router.DELETE("/admin/verification-code/:id", services.DeleteVerificationCode)
}

func GetUserContacts(router *gin.Engine) {
	router.GET("/admin/user/contacts", services.GetUserContacts)
}

func GetAllManagersContacts(router *gin.Engine) {
	router.GET("/admin/managers/contacts", services.GetAllManagersContacts)
}

func GetAllManagers(router *gin.Engine) {
	router.GET("/admin/managers", services.GetAllManagers)
}
