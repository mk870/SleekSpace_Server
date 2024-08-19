package insights

import (
	"SleekSpace/middleware"
	insightsService "SleekSpace/services/property/insights"

	"github.com/gin-gonic/gin"
)

var groupRouteName string = "/property/"

func GetPropertyInsightsById(router *gin.Engine) {
	router.GET(groupRouteName+"insights/:id", middleware.AuthValidator, insightsService.GetPropertyInsightsById)
}

func GetPropertyInsightsByPropertyId(router *gin.Engine) {
	router.GET(groupRouteName+"insights/property/:id", middleware.AuthValidator, insightsService.GetPropertyInsightsByPropertyId)
}

func UpdatePropertyInsights(router *gin.Engine) {
	router.PUT(groupRouteName+"insights/:id", middleware.AuthValidator, insightsService.UpdatePropertyInsights)
}
