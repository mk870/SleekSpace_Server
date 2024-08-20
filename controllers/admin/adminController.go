package controllers

import (
	adminService "SleekSpace/services/admin"

	"github.com/gin-gonic/gin"
)

func GetAllUsersLocations(router *gin.Engine) {
	router.GET("/admin/users/location", adminService.GetAllUsersLocations)
}

func GetAllManagersProfilePictures(router *gin.Engine) {
	router.GET("/admin/managers/profile-pictures", adminService.GetAllManagersProfilePictures)
}

func ApiSim(router *gin.Engine) {
	router.GET("/admin/sim", adminService.GetInfo)
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

func GetAllStands(router *gin.Engine) {
	router.GET("/admin/property/stands", adminService.GetAllStands)
}

func GetAllLandProperties(router *gin.Engine) {
	router.GET("/admin/property/lands", adminService.GetAllLandProperties)
}

func GetAllCommercialRentalProperties(router *gin.Engine) {
	router.GET("/admin/property/commercial/rentals", adminService.GetAllCommercialRentalProperties)
}

func GetAllCommercialPropertiesForSale(router *gin.Engine) {
	router.GET("/admin/property/commercial/onsale", adminService.GetAllCommercialForSaleProperties)
}

func GetAllResidentialPropertiesForSale(router *gin.Engine) {
	router.GET("/admin/property/residential/onsale", adminService.GetAllResidentialForSaleProperties)
}

func GetAllResidentialRentalProperties(router *gin.Engine) {
	router.GET("/admin/property/residential/rentals", adminService.GetAllResidentialRentalProperties)
}

func GetAllPropertiesImagesOrVideos(router *gin.Engine) {
	router.GET("/admin/property/media", adminService.GetAllPropertiesImagesOrVideos)
}

func GetAllPropertiesLocations(router *gin.Engine) {
	router.GET("/admin/property/location", adminService.GetAllPropertiesLocation)
}

func GetAllPropertiesInsights(router *gin.Engine) {
	router.GET("/admin/property/insights", adminService.GetAllPropertiesInsights)
}
