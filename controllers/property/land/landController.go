package land

import (
	"SleekSpace/middleware"
	landService "SleekSpace/services/property/land"

	"github.com/gin-gonic/gin"
)

var landRoute string = "/property/land"

func CreateLandForSale(router *gin.Engine) {
	router.POST(landRoute, middleware.AuthValidator, landService.CreateLandPropertyForSale)
}

func GetLandById(router *gin.Engine) {
	router.GET(landRoute+"/:id", middleware.AuthValidator, landService.GetLandPropertyById)
}

func UpdateLandDetails(router *gin.Engine) {
	router.PUT(landRoute+"/:id", middleware.AuthValidator, landService.UpdateLandPropertyDetails)
}

func DeleteLandById(router *gin.Engine) {
	router.DELETE(landRoute+"/:id", middleware.AuthValidator, landService.DeleteLandPropertyById)
}
