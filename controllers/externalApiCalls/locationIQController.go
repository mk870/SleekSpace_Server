package controllers

import (
	"SleekSpace/httpServices/location"

	"github.com/gin-gonic/gin"
)

func LocationAutoComplete(router *gin.Engine) {
	router.POST("/location/autocomplete", location.LocationAutoComplete)
}

func LocationReverseGeoCoding(router *gin.Engine) {
	router.POST("/location/reverse-geocoding", location.LocationReverseGeoCoding)
}
