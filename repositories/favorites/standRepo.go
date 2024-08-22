package favorites

import (
	"SleekSpace/db"
	propertyModels "SleekSpace/models/property"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFavoriteStandProperties(standPropertiesIds []int) []propertyModels.PropertyStand {
	var properties = []propertyModels.PropertyStand{}
	result := db.DB.Where("id IN ?", standPropertiesIds).Preload(clause.Associations).Find(&properties)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return properties
}
