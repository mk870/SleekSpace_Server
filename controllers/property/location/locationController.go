package location

import (
	"SleekSpace/middleware"
	locationService "SleekSpace/services/property/location"

	"github.com/gin-gonic/gin"
)

var groupRouteName string = "/property/"

func GetPropertyLocationById(router *gin.Engine) {
	router.GET(groupRouteName+"location/:id", middleware.AuthValidator, locationService.GetPropertyLocationById)
}

func UpdatePropertyLocation(router *gin.Engine) {
	router.PUT(groupRouteName+"location/:id", middleware.AuthValidator, locationService.UpdatePropertyLocation)
}
