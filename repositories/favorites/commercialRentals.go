package favorites

import (
	"SleekSpace/db"
	propertyModels "SleekSpace/models/property"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFavoriteCommercialRentalProperties(commercialRentalPropertiesIds []int) []propertyModels.CommercialRentalProperty {
	var properties = []propertyModels.CommercialRentalProperty{}
	result := db.DB.Where("id IN ?", commercialRentalPropertiesIds).Preload(clause.Associations).Find(&properties)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return properties
}
