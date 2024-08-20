package commercial

import (
	"SleekSpace/middleware"
	commercialService "SleekSpace/services/property/commercial"

	"github.com/gin-gonic/gin"
)

var rentalRoute string = "/property/commercial/rentals"

func CreateCommercialRentalProperty(router *gin.Engine) {
	router.POST(rentalRoute, middleware.AuthValidator, commercialService.CreateCommercialRentalProperty)
}

func GetCommercialRentalPropertyById(router *gin.Engine) {
	router.GET(rentalRoute+"/:id", middleware.AuthValidator, commercialService.GetCommercialRentalPropertyId)
}

func UpdateCommercialRentalPropertyDetails(router *gin.Engine) {
	router.PUT(rentalRoute+"/:id", middleware.AuthValidator, commercialService.UpdateCommercialRentalPropertyDetails)
}

func DeleteCommercialRentalPropertyId(router *gin.Engine) {
	router.DELETE(rentalRoute+"/:id", middleware.AuthValidator, commercialService.DeleteCommercialRentalPropertyById)
}
