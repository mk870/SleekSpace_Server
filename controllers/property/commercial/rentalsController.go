package commercial

import (
	"SleekSpace/middleware"
	commercialService "SleekSpace/services/property/commercial"

	"github.com/gin-gonic/gin"
)

func CommercialRentalPropertyForSaleRoutes(router *gin.Engine) {
	routes := router.Group("/property/commercial/rentals")
	{
		routes.POST("", middleware.AuthValidator, commercialService.CreateCommercialRentalProperty)
		routes.GET("", commercialService.GetAllCommercialRentalProperties)
		routes.GET("/:id", commercialService.GetCommercialRentalPropertyId)
		routes.PUT("/:id", middleware.AuthValidator, commercialService.UpdateCommercialRentalPropertyDetails)
		routes.DELETE("/:id", middleware.AuthValidator, commercialService.DeleteCommercialRentalPropertyById)
	}
}
