package stand

import (
	"SleekSpace/db"
	managerModels "SleekSpace/models/manager"
	propertyModels "SleekSpace/models/property"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateStandForSale(manager *managerModels.Manager, stand *propertyModels.PropertyStand) bool {
	err := db.DB.Model(manager).Association("PropertyStand").Append(stand)
	if err != nil {
		println(err.Error())
	}
	return true
}

func GetStandById(id string) *propertyModels.PropertyStand {
	var stand propertyModels.PropertyStand
	result := db.DB.First(&stand, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &stand
}

func GetStandWithAllAssociationsById(id string) *propertyModels.PropertyStand {
	var stand propertyModels.PropertyStand
	result := db.DB.Preload(clause.Associations).First(&stand, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &stand
}

func GetManagerStandsByManagerId(managerId string) []propertyModels.PropertyStand {
	var manager = managerModels.Manager{}
	result := db.DB.Preload("PropertyStand").First(&manager, managerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return manager.PropertyStand
}

func GetAllStands() []propertyModels.PropertyStand {
	var stands = []propertyModels.PropertyStand{}
	err := db.DB.Find(&stands)
	if err != nil {
		println(err.Error, err.Name())
	}
	return stands
}

func UpdateStand(update *propertyModels.PropertyStand) bool {
	db.DB.Save(update)
	return true
}

func DeleteStandById(id string) bool {
	manager := GetStandById(id)
	if manager == nil {
		return false
	}
	db.DB.Select(clause.Associations).Unscoped().Delete(&manager)
	return true
}
