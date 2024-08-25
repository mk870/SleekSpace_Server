package commercial

import (
	"SleekSpace/middleware"
	commercialService "SleekSpace/services/property/commercial"

	"github.com/gin-gonic/gin"
)

func CommercialPropertyForSaleRoutes(router *gin.Engine) {
	routes := router.Group("/property/commercial/onsale")
	{
		routes.POST("", middleware.AuthValidator, commercialService.CreateCommercialPropertyForSale)
		routes.GET("", commercialService.GetAllCommercialForSaleProperties)
		routes.GET("/:id", middleware.AuthValidator, commercialService.GetCommercialPropertyForSaleById)
		routes.PUT("/:id", middleware.AuthValidator, commercialService.UpdateCommercialPropertyForSaleDetails)
		routes.DELETE("/:id", middleware.AuthValidator, commercialService.DeleteCommercialPropertyForSaleById)
	}
}
