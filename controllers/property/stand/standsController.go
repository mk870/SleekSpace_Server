package stand

import (
	"SleekSpace/middleware"
	standService "SleekSpace/services/property/stand"

	"github.com/gin-gonic/gin"
)

var groupRouteName string = "/property/"

func CreateStandForSale(router *gin.Engine) {
	router.POST(groupRouteName+"stands", middleware.AuthValidator, standService.CreateStandForSale)
}

func GetStandById(router *gin.Engine) {
	router.GET(groupRouteName+"stands/:id", middleware.AuthValidator, standService.GetStandById)
}

func UpdateStandDetails(router *gin.Engine) {
	router.PUT(groupRouteName+"stands/:id", middleware.AuthValidator, standService.UpdateStandDetails)
}

func DeleteStandById(router *gin.Engine) {
	router.DELETE(groupRouteName+"stands/:id", middleware.AuthValidator, standService.DeleteStandById)
}
