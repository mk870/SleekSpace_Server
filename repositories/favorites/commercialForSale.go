package favorites

import (
	"SleekSpace/db"
	managerModels "SleekSpace/models/manager"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetFavoriteCommercialPropertiesForSale(commercialForSalePropertiesIds []int) []managerModels.CommercialForSaleProperty {
	var properties = []managerModels.CommercialForSaleProperty{}
	result := db.DB.Where("id IN ?", commercialForSalePropertiesIds).Preload(clause.Associations).Preload("Manager.ProfilePicture").Preload("Manager.ManagerContactNumbers").Find(&properties)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return properties
}
