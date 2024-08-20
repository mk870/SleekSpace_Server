package residential

import (
	"SleekSpace/middleware"
	residentialService "SleekSpace/services/property/residential"

	"github.com/gin-gonic/gin"
)

var rentalRoute string = "/property/residential/rentals"

func CreateResidentialRentalProperty(router *gin.Engine) {
	router.POST(rentalRoute, middleware.AuthValidator, residentialService.CreateResidentialRentalProperty)
}

func GetResidentialRentalPropertyById(router *gin.Engine) {
	router.GET(rentalRoute+"/:id", middleware.AuthValidator, residentialService.GetResidentialRentalPropertyId)
}

func UpdateResidentialRentalPropertyDetails(router *gin.Engine) {
	router.PUT(rentalRoute+"/:id", middleware.AuthValidator, residentialService.UpdateResidentialRentalPropertyDetails)
}

func DeleteResidentialRentalPropertyById(router *gin.Engine) {
	router.DELETE(rentalRoute+"/:id", middleware.AuthValidator, residentialService.DeleteResidentialRentalPropertyById)
}
