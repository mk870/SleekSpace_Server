package favorites

import (
	"SleekSpace/db"
	propertyModels "SleekSpace/models/property"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFavoriteLandProperties(landPropertiesIds []int) []propertyModels.LandForSaleProperty {
	var properties = []propertyModels.LandForSaleProperty{}
	result := db.DB.Where("id IN ?", landPropertiesIds).Preload(clause.Associations).Find(&properties)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return properties
}
