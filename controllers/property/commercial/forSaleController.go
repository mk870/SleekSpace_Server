package commercial

import (
	"SleekSpace/middleware"
	commercialService "SleekSpace/services/property/commercial"

	"github.com/gin-gonic/gin"
)

var forSaleRoute string = "/property/commercial/onsale"

func CreateCommercialPropertyForSale(router *gin.Engine) {
	router.POST(forSaleRoute, middleware.AuthValidator, commercialService.CreateCommercialPropertyForSale)
}

func GetCommercialPropertyForSaleById(router *gin.Engine) {
	router.GET(forSaleRoute+"/:id", middleware.AuthValidator, commercialService.GetCommercialPropertyForSaleById)
}

func UpdateCommercialPropertyForSaleDetails(router *gin.Engine) {
	router.PUT(forSaleRoute+"/:id", middleware.AuthValidator, commercialService.UpdateCommercialPropertyForSaleDetails)
}

func DeleteCommercialPropertyForSaleById(router *gin.Engine) {
	router.DELETE(forSaleRoute+"/:id", middleware.AuthValidator, commercialService.DeleteCommercialPropertyForSaleById)
}
