package favorites

import (
	"SleekSpace/db"
	propertyModels "SleekSpace/models/property"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFavoriteCommercialPropertiesForSale(commercialForSalePropertiesIds []int) []propertyModels.CommercialForSaleProperty {
	var properties = []propertyModels.CommercialForSaleProperty{}
	result := db.DB.Where("id IN ?", commercialForSalePropertiesIds).Preload(clause.Associations).Find(&properties)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return properties
}
