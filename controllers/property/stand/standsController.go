package stand

import (
	"SleekSpace/middleware"
	standService "SleekSpace/services/property/stand"

	"github.com/gin-gonic/gin"
)

var standRoute string = "/property/stand"

func CreateStandForSale(router *gin.Engine) {
	router.POST(standRoute, middleware.AuthValidator, standService.CreateStandForSale)
}

func GetStandById(router *gin.Engine) {
	router.GET(standRoute+"/:id", middleware.AuthValidator, standService.GetStandById)
}

func UpdateStandDetails(router *gin.Engine) {
	router.PUT(standRoute+"/:id", middleware.AuthValidator, standService.UpdateStandDetails)
}

func DeleteStandById(router *gin.Engine) {
	router.DELETE(standRoute+"/:id", middleware.AuthValidator, standService.DeleteStandById)
}
