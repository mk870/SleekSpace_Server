package commercial

import (
	"SleekSpace/db"
	managerModels "SleekSpace/models/manager"
	propertyModels "SleekSpace/models/property"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateCommercialRentalProperty(property *propertyModels.CommercialRentalProperty) bool {
	err := db.DB.Create(property)
	if err != nil {
		println(err)
	}
	return true
}

func GetCommercialRentalPropertyById(id string) *propertyModels.CommercialRentalProperty {
	var property propertyModels.CommercialRentalProperty
	result := db.DB.First(&property, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetCommercialRentalPropertyWithAllAssociationsById(id string) *propertyModels.CommercialRentalProperty {
	var property propertyModels.CommercialRentalProperty
	result := db.DB.Preload(clause.Associations).First(&property, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetCommercialRentalPropertyWithAllAssociationsByUniqueId(uniqueId string) *propertyModels.CommercialRentalProperty {
	var property propertyModels.CommercialRentalProperty
	result := db.DB.Where("unique_id= ?", uniqueId).Preload(clause.Associations).First(&property)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetManagerCommercialRentalPropertiesByManagerId(managerId string) []propertyModels.CommercialRentalProperty {
	var manager = managerModels.Manager{}
	result := db.DB.Preload("CommercialRentalProperty").First(&manager, managerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return manager.CommercialRentalProperty
}

func GetAllCommercialRentalProperties() []propertyModels.CommercialRentalProperty {
	var properties = []propertyModels.CommercialRentalProperty{}
	err := db.DB.Preload(clause.Associations).Find(&properties)
	if err != nil {
		println(err.Error, err.Name())
	}
	return properties
}

func UpdateCommercialRentalProperty(update *propertyModels.CommercialRentalProperty) bool {
	db.DB.Save(update)
	return true
}

func DeleteCommercialRentalPropertyById(id string) bool {
	property := GetCommercialRentalPropertyById(id)
	if property == nil {
		return false
	}
	db.DB.Select(clause.Associations).Unscoped().Delete(&property)
	return true
}