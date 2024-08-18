package insights

import (
	"SleekSpace/db"
	propertyModels "SleekSpace/models/property"
)

func GetAllPropertyInsights() []propertyModels.PropertyInsights {
	var insightsList = []propertyModels.PropertyInsights{}
	err := db.DB.Find(&insightsList)
	if err != nil {
		println(err.Error, err.Name())
	}
	return insightsList
}

func GetPropertyInsightsById(propertyInsightsId int) propertyModels.PropertyInsights {
	var insights = propertyModels.PropertyInsights{}
	db.DB.First(&insights, propertyInsightsId)
	return insights
}

func UpdatePropertyInsights(propertyInsightsUpdate *propertyModels.PropertyInsights) bool {
	db.DB.Save(propertyInsightsUpdate)
	return true
}

func DeletePropertyInsights(propertyInsightsId string) bool {
	db.DB.Unscoped().Delete(&propertyModels.PropertyInsights{}, propertyInsightsId)
	return true
}
