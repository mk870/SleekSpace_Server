package land

import (
	"SleekSpace/middleware"
	landService "SleekSpace/services/property/land"

	"github.com/gin-gonic/gin"
)

func LandPropertyRoutes(router *gin.Engine) {
	routes := router.Group("/property/land")
	{
		routes.POST("", middleware.AuthValidator, landService.CreateLandPropertyForSale)
		routes.GET("", landService.GetAllLandProperties)
		routes.GET("/:id", middleware.AuthValidator, landService.GetLandPropertyById)
		routes.PUT("/:id", middleware.AuthValidator, landService.UpdateLandPropertyDetails)
		routes.DELETE("/:id", middleware.AuthValidator, landService.DeleteLandPropertyById)
	}
}
