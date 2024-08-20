package controllers

import (
	"SleekSpace/middleware"
	managerService "SleekSpace/services/manager"
	commercialService "SleekSpace/services/property/commercial"
	landService "SleekSpace/services/property/land"
	residentialService "SleekSpace/services/property/residential"
	standService "SleekSpace/services/property/stand"

	"github.com/gin-gonic/gin"
)

func CreateManager(router *gin.Engine) {
	router.POST("/manager", middleware.AuthValidator, managerService.CreateManager)
}

func GetManagerByUserId(router *gin.Engine) {
	router.GET("/manager/user/:userId", middleware.AuthValidator, managerService.GetManagerByUserId)
}

func GetManagerById(router *gin.Engine) {
	router.GET("/manager/:id", middleware.AuthValidator, managerService.GetManagerByManagerId)
}

func GetManagerStandsByManagerId(router *gin.Engine) {
	router.GET("/manager/stands/:id", middleware.AuthValidator, standService.GetManagerStandsByManagerId)
}

func GetManagerLandPropertiesByManagerId(router *gin.Engine) {
	router.GET("/manager/lands/:id", middleware.AuthValidator, landService.GetManagerLandPropertiesByManagerId)
}

func GetManagerCommercialRentalPropertiesByManagerId(router *gin.Engine) {
	router.GET("/manager/commercial/rentals/:id", middleware.AuthValidator, commercialService.GetManagerCommercialRentalPropertiesByManagerId)
}

func GetManagerCommercialPropertiesForSaleByManagerId(router *gin.Engine) {
	router.GET("/manager/commercial/onsale/:id", middleware.AuthValidator, commercialService.GetManagerCommercialPropertiesForSaleByManagerId)
}

func GetManagerResidentialRentalPropertiesByManagerId(router *gin.Engine) {
	router.GET("/manager/residential/rentals/:id", middleware.AuthValidator, residentialService.GetManagerResidentialRentalPropertiesByManagerId)
}

func GetManagerResidentialPropertiesForSaleByManagerId(router *gin.Engine) {
	router.GET("/manager/residential/onsale/:id", middleware.AuthValidator, residentialService.GetManagerResidentialPropertiesForSaleByManagerId)
}

func UpdateManager(router *gin.Engine) {
	router.PUT("/manager/:id", middleware.AuthValidator, managerService.UpdateManagerEmailAndName)
}

func DeleteManager(router *gin.Engine) {
	router.DELETE("/manager/:id", middleware.AuthValidator, managerService.DeleteManager)
}
