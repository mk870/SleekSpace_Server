package media

import (
	"SleekSpace/middleware"
	mediaService "SleekSpace/services/property/media"

	"github.com/gin-gonic/gin"
)

var groupRouteName string = "/property/"

func CreatePropertyImageOrVideo(router *gin.Engine) {
	router.POST(groupRouteName+"media", middleware.AuthValidator, mediaService.CreatePropertyImageOrVideoWithPropertyId)
}

func GetPropertyImageOrVideoById(router *gin.Engine) {
	router.GET(groupRouteName+"media/:id", middleware.AuthValidator, mediaService.GetPropertyImageOrVideoById)
}

func UpdatePropertyImageOrVideo(router *gin.Engine) {
	router.PUT(groupRouteName+"media/:id", middleware.AuthValidator, mediaService.UpdatePropertyImageOrVideo)
}

func DeletePropertyImageOrVideo(router *gin.Engine) {
	router.DELETE(groupRouteName+"media/:id", middleware.AuthValidator, mediaService.DeletePropertyImageOrVideo)
}
