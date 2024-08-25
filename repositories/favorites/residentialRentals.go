package favorites

import (
	"SleekSpace/db"
	managerModels "SleekSpace/models/manager"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFavoriteResidentialRentalProperties(residentialRentalPropertiesIds []int) []managerModels.ResidentialRentalProperty {
	var properties = []managerModels.ResidentialRentalProperty{}
	result := db.DB.Where("id IN ?", residentialRentalPropertiesIds).Preload(clause.Associations).Preload("Manager.ProfilePicture").Preload("Manager.ManagerContactNumbers").Find(&properties)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return properties
}
