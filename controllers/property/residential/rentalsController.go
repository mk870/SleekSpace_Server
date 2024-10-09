package residential

import (
	"SleekSpace/middleware"
	residentialService "SleekSpace/services/property/residential"

	"github.com/gin-gonic/gin"
)

func ResidentialRentalPropertyRoutes(router *gin.Engine) {
	routes := router.Group("/property/residential/rentals")
	{
		routes.POST("", middleware.AuthValidator, residentialService.CreateResidentialRentalProperty)
		routes.GET("", residentialService.GetAllResidentialRentalProperties)
		routes.GET("/:id", residentialService.GetResidentialRentalPropertyId)
		routes.PUT("/:id", middleware.AuthValidator, residentialService.UpdateResidentialRentalPropertyDetails)
		routes.DELETE("/:id", middleware.AuthValidator, residentialService.DeleteResidentialRentalPropertyById)
	}
}
