package favorites

import (
	"SleekSpace/db"
	managerModels "SleekSpace/models/manager"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFavoriteResidentialPropertiesForSale(residentialForSalePropertiesIds []int) []managerModels.ResidentialPropertyForSale {
	var properties = []managerModels.ResidentialPropertyForSale{}
	result := db.DB.Where("id IN ?", residentialForSalePropertiesIds).Preload(clause.Associations).Preload("Manager.ProfilePicture").Preload("Manager.ManagerContactNumbers").Find(&properties)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return properties
}
