package controllers

import (
	"SleekSpace/httpServices/location"
	"SleekSpace/middleware"
	"SleekSpace/services"

	"github.com/gin-gonic/gin"
)

func CreateLocation(router *gin.Engine) {
	router.POST("/location", middleware.AuthValidator, services.CreateLocation)
}

func UpdateLocation(router *gin.Engine) {
	router.PUT("/location", middleware.AuthValidator, services.UpdateLocation)
}

func LocationAutoComplete(router *gin.Engine) {
	router.POST("/location/autocomplete", location.LocationAutoComplete)
}

func LocationReverseGeoCoding(router *gin.Engine) {
	router.POST("/location/reverse-geocoding", location.LocationReverseGeoCoding)
}
