package residential

import (
	"SleekSpace/db"
	managerModels "SleekSpace/models/manager"
	propertyModels "SleekSpace/models/property"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateResidentialRentalProperty(property *propertyModels.ResidentialRentalProperty) bool {
	err := db.DB.Create(property)
	if err != nil {
		println(err)
	}
	return true
}

func GetResidentialRentalPropertyById(id string) *propertyModels.ResidentialRentalProperty {
	var property propertyModels.ResidentialRentalProperty
	result := db.DB.First(&property, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetResidentialRentalPropertyWithAllAssociationsById(id string) *propertyModels.ResidentialRentalProperty {
	var property propertyModels.ResidentialRentalProperty
	result := db.DB.Preload(clause.Associations).First(&property, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetManagerResidentialRentalPropertiesByManagerId(managerId string) []propertyModels.ResidentialRentalProperty {
	var manager = managerModels.Manager{}
	result := db.DB.Preload("ResidentialRentalProperty").First(&manager, managerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return manager.ResidentialRentalProperty
}

func GetAllResidentialRentalProperties() []propertyModels.ResidentialRentalProperty {
	var properties = []propertyModels.ResidentialRentalProperty{}
	err := db.DB.Preload(clause.Associations).Find(&properties)
	if err != nil {
		println(err.Error, err.Name())
	}
	return properties
}

func UpdateResidentialRentalProperty(update *propertyModels.ResidentialRentalProperty) bool {
	db.DB.Save(update)
	return true
}

func DeleteResidentialRentalPropertyById(id string) bool {
	property := GetResidentialRentalPropertyById(id)
	if property == nil {
		return false
	}
	db.DB.Select(clause.Associations).Unscoped().Delete(&property)
	return true
}
