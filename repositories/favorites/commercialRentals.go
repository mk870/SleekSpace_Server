package favorites

import (
	"SleekSpace/db"
	managerModels "SleekSpace/models/manager"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFavoriteCommercialRentalProperties(commercialRentalPropertiesIds []int) []managerModels.CommercialRentalProperty {
	var properties = []managerModels.CommercialRentalProperty{}
	result := db.DB.Where("id IN ?", commercialRentalPropertiesIds).Preload(clause.Associations).Preload("Manager.ProfilePicture").Preload("Manager.ManagerContactNumbers").Find(&properties)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return properties
}
