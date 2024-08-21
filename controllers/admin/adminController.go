package controllers

import (
	adminService "SleekSpace/services/admin"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine) {
	routes := router.Group("/admin")
	{
		routes.GET("/users", adminService.GetAllUsers)
		routes.GET("/users/location", adminService.GetAllUsersLocations)
		routes.GET("/managers/profile-pictures", adminService.GetAllManagersProfilePictures)
		routes.GET("/sim", adminService.GetInfo)
		routes.GET("/location/:id", adminService.GetLocationById)
		routes.GET("/verification-code/:id", adminService.GetVerificationCodeById)
		routes.GET("/verification-codes", adminService.GetAllVerificationCodes)
		routes.DELETE("/verification-code/:id", adminService.DeleteVerificationCode)
		routes.GET("/user/contacts", adminService.GetUserContacts)
		routes.GET("/managers/contacts", adminService.GetAllManagersContacts)
		routes.GET("/managers", adminService.GetAllManagers)
		routes.GET("/property/stands", adminService.GetAllStands)
		routes.GET("/property/lands", adminService.GetAllLandProperties)
		routes.GET("/property/commercial/rentals", adminService.GetAllCommercialRentalProperties)
		routes.GET("/property/commercial/onsale", adminService.GetAllCommercialForSaleProperties)
		routes.GET("/property/residential/onsale", adminService.GetAllResidentialForSaleProperties)
		routes.GET("/property/residential/rentals", adminService.GetAllResidentialRentalProperties)
		routes.GET("/property/media", adminService.GetAllPropertiesImagesOrVideos)
		routes.GET("/property/location", adminService.GetAllPropertiesLocation)
		routes.GET("/property/insights", adminService.GetAllPropertiesInsights)
	}
}
