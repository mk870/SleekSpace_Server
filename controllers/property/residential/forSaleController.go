package residential

import (
	"SleekSpace/middleware"
	residentialService "SleekSpace/services/property/residential"

	"github.com/gin-gonic/gin"
)

func ResidentialPropertyForSaleRoutes(router *gin.Engine) {
	routes := router.Group("/property/residential/onsale")
	{
		routes.POST("", middleware.AuthValidator, residentialService.CreateResidentialPropertyForSale)
		routes.GET("", residentialService.GetAllResidentialForSaleProperties)
		routes.GET("/:id", middleware.AuthValidator, residentialService.GetResidentialPropertyForSaleById)
		routes.DELETE("/:id", middleware.AuthValidator, residentialService.DeleteResidentialPropertyForSaleById)
		routes.PUT("/:id", middleware.AuthValidator, residentialService.UpdateResidentialPropertyForSaleDetails)
	}
}
