package favorites

import (
	"SleekSpace/db"
	propertyModels "SleekSpace/models/property"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFavoriteResidentialRentalProperties(residentialRentalPropertiesIds []int) []propertyModels.ResidentialRentalProperty {
	var properties = []propertyModels.ResidentialRentalProperty{}
	result := db.DB.Where("id IN ?", residentialRentalPropertiesIds).Preload(clause.Associations).Find(&properties)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return properties
}
