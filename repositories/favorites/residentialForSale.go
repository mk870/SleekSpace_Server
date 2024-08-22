package favorites

import (
	"SleekSpace/db"
	propertyModels "SleekSpace/models/property"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFavoriteResidentialPropertiesForSale(residentialForSalePropertiesIds []int) []propertyModels.ResidentialPropertyForSale {
	var properties = []propertyModels.ResidentialPropertyForSale{}
	result := db.DB.Where("id IN ?", residentialForSalePropertiesIds).Preload(clause.Associations).Find(&properties)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return properties
}
