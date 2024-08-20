package commercial

import (
	"SleekSpace/db"
	managerModels "SleekSpace/models/manager"
	propertyModels "SleekSpace/models/property"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateCommercialPropertyForSale(property *propertyModels.CommercialForSaleProperty) bool {
	err := db.DB.Create(property)
	if err != nil {
		println(err)
	}
	return true
}

func GetCommercialPropertyForSaleById(id string) *propertyModels.CommercialForSaleProperty {
	var property propertyModels.CommercialForSaleProperty
	result := db.DB.First(&property, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetCommercialPropertyForSaleWithAllAssociationsById(id string) *propertyModels.CommercialForSaleProperty {
	var property propertyModels.CommercialForSaleProperty
	result := db.DB.Preload(clause.Associations).First(&property, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &property
}

func GetManagerCommercialPropertiesForSaleByManagerId(managerId string) []propertyModels.CommercialForSaleProperty {
	var manager = managerModels.Manager{}
	result := db.DB.Preload("CommercialForSaleProperty").First(&manager, managerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return manager.CommercialForSaleProperty
}

func GetAllCommercialPropertiesForSale() []propertyModels.CommercialForSaleProperty {
	var properties = []propertyModels.CommercialForSaleProperty{}
	err := db.DB.Preload(clause.Associations).Find(&properties)
	if err != nil {
		println(err.Error, err.Name())
	}
	return properties
}

func UpdateCommercialPropertyForSale(update *propertyModels.CommercialForSaleProperty) bool {
	db.DB.Save(update)
	return true
}

func DeleteCommercialPropertyForSaleById(id string) bool {
	property := GetCommercialPropertyForSaleById(id)
	if property == nil {
		return false
	}
	db.DB.Select(clause.Associations).Unscoped().Delete(&property)
	return true
}
