package land

import (
	"SleekSpace/db"
	managerModels "SleekSpace/models/manager"
	propertyModels "SleekSpace/models/property"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateLandPropertyForSale(land *propertyModels.LandForSaleProperty) bool {
	err := db.DB.Create(land)
	if err != nil {
		println(err)
	}
	return true
}

func GetLandPropertyForSaleById(id string) *propertyModels.LandForSaleProperty {
	var land propertyModels.LandForSaleProperty
	result := db.DB.First(&land, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &land
}

func GetLandPropertyForSaleWithAllAssociationsById(id string) *propertyModels.LandForSaleProperty {
	var land propertyModels.LandForSaleProperty
	result := db.DB.Preload(clause.Associations).First(&land, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &land
}

func GetManagerLandPropertiesForSaleByManagerId(managerId string) []propertyModels.LandForSaleProperty {
	var manager = managerModels.Manager{}
	result := db.DB.Preload("LandForSaleProperty").First(&manager, managerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return manager.LandForSaleProperty
}

func GetAllLandPropertiesForSale() []propertyModels.LandForSaleProperty {
	var properties = []propertyModels.LandForSaleProperty{}
	err := db.DB.Preload(clause.Associations).Find(&properties)
	if err != nil {
		println(err.Error, err.Name())
	}
	return properties
}

func UpdateLandPropertyForSale(update *propertyModels.LandForSaleProperty) bool {
	db.DB.Save(update)
	return true
}

func DeleteLandPropertyForSaleById(id string) bool {
	land := GetLandPropertyForSaleById(id)
	if land == nil {
		return false
	}
	db.DB.Select(clause.Associations).Unscoped().Delete(&land)
	return true
}
