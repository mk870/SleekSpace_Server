package favorites

import (
	"SleekSpace/db"
	managerModels "SleekSpace/models/manager"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFavoriteLandProperties(landPropertiesIds []int) []managerModels.LandForSaleProperty {
	var properties = []managerModels.LandForSaleProperty{}
	result := db.DB.Where("id IN ?", landPropertiesIds).Preload(clause.Associations).Preload("Manager.ProfilePicture").Preload("Manager.ManagerContactNumbers").Find(&properties)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return properties
}
