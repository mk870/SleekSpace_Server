package controllers

import (
	adminService "SleekSpace/services/admin"

	"github.com/gin-gonic/gin"
)

func GetAllUsersLocations(router *gin.Engine) {
	router.GET("/admin/users/location", adminService.GetAllUsersLocations)
}

func GetLocationById(router *gin.Engine) {
	router.GET("/admin/location/:id", adminService.GetLocationById)
}

func GetVerificationCodeById(router *gin.Engine) {
	router.GET("/admin/verification-code/:id", adminService.GetVerificationCodeById)
}

func GetAllVerificationCodes(router *gin.Engine) {
	router.GET("/admin/verification-codes", adminService.GetAllVerificationCodes)
}

func DeleteVerificationCode(router *gin.Engine) {
	router.DELETE("/admin/verification-code/:id", adminService.DeleteVerificationCode)
}

func GetUserContacts(router *gin.Engine) {
	router.GET("/admin/user/contacts", adminService.GetUserContacts)
}

func GetAllManagersContacts(router *gin.Engine) {
	router.GET("/admin/managers/contacts", adminService.GetAllManagersContacts)
}

func GetAllManagers(router *gin.Engine) {
	router.GET("/admin/managers", adminService.GetAllManagers)
}
