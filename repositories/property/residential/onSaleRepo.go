package residential

import (
	"SleekSpace/db"
	managerModels "SleekSpace/models/manager"
	propertyModels "SleekSpace/models/property"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateResidentialPropertyForSale(property *propertyModels.ResidentialPropertyForSale) bool {
	err := db.DB.Create(property)
	if err != nil {
		println(err)
	}
	return true
}

func GetResidentialPropertyForSaleById(id string) *propertyModels.ResidentialPropertyForSale {
	var property propertyModels.ResidentialPropertyForSale
	result := db.DB.First(&property, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetResidentialPropertyForSaleWithAllAssociationsById(id string) *propertyModels.ResidentialPropertyForSale {
	var property propertyModels.ResidentialPropertyForSale
	result := db.DB.Preload(clause.Associations).First(&property, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetResidentialPropertyForSaleWithAllAssociationsByUniqueId(uniqueId string) *propertyModels.ResidentialPropertyForSale {
	var property propertyModels.ResidentialPropertyForSale
	result := db.DB.Where("unique_id= ?", uniqueId).Preload(clause.Associations).First(&property)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetManagerResidentialPropertiesForSaleByManagerId(managerId string) []propertyModels.ResidentialPropertyForSale {
	var manager = managerModels.Manager{}
	result := db.DB.Preload("ResidentialPropertyForSale").First(&manager, managerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return manager.ResidentialPropertyForSale
}

func GetAllResidentialPropertiesForSale() []propertyModels.ResidentialPropertyForSale {
	var properties = []propertyModels.ResidentialPropertyForSale{}
	err := db.DB.Preload(clause.Associations).Find(&properties)
	if err != nil {
		println(err.Error, err.Name())
	}
	return properties
}

func UpdateResidentialPropertyForSale(update *propertyModels.ResidentialPropertyForSale) bool {
	db.DB.Save(update)
	return true
}

func DeleteResidentialPropertyForSaleById(id string) bool {
	property := GetResidentialPropertyForSaleById(id)
	if property == nil {
		return false
	}
	db.DB.Select(clause.Associations).Unscoped().Delete(&property)
	return true
}
