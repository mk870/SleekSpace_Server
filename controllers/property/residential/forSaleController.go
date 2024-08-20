package residential

import (
	"SleekSpace/middleware"
	residentialService "SleekSpace/services/property/residential"

	"github.com/gin-gonic/gin"
)

var forSaleRoute string = "/property/residential/onsale"

func CreateResidentialPropertyForSale(router *gin.Engine) {
	router.POST(forSaleRoute, middleware.AuthValidator, residentialService.CreateResidentialPropertyForSale)
}

func GetResidentialPropertyForSaleById(router *gin.Engine) {
	router.GET(forSaleRoute+"/:id", middleware.AuthValidator, residentialService.GetResidentialPropertyForSaleById)
}

func UpdateResidentialPropertyForSaleDetails(router *gin.Engine) {
	router.PUT(forSaleRoute+"/:id", middleware.AuthValidator, residentialService.UpdateResidentialPropertyForSaleDetails)
}

func DeleteResidentialPropertyForSaleById(router *gin.Engine) {
	router.DELETE(forSaleRoute+"/:id", middleware.AuthValidator, residentialService.DeleteResidentialPropertyForSaleById)
}
