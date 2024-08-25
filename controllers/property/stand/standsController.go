package stand

import (
	"SleekSpace/middleware"
	standService "SleekSpace/services/property/stand"

	"github.com/gin-gonic/gin"
)

func StandsRoutes(router *gin.Engine) {
	routes := router.Group("/property/stand")
	{
		routes.POST("", middleware.AuthValidator, standService.CreateStandForSale)
		routes.GET("", standService.GetAllStands)
		routes.GET("/:id", middleware.AuthValidator, standService.GetStandById)
		routes.PUT("/:id", middleware.AuthValidator, standService.UpdateStandDetails)
		routes.DELETE("/:id", middleware.AuthValidator, standService.DeleteStandById)
	}
}
