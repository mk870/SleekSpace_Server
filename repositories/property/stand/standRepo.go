package stand

import (
	"SleekSpace/db"
	managerModels "SleekSpace/models/manager"
	propertyModels "SleekSpace/models/property"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateStandForSale(stand *propertyModels.PropertyStand) bool {
	err := db.DB.Create(stand)
	if err != nil {
		println(err)
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

func GetStandWithAllAssociationsByUniqueId(uniqueId string) *propertyModels.PropertyStand {
	var stand propertyModels.PropertyStand
	result := db.DB.Where("unique_id= ?", uniqueId).Preload(clause.Associations).First(&stand)
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
	err := db.DB.Preload(clause.Associations).Find(&stands)
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
	stand := GetStandById(id)
	if stand == nil {
		return false
	}
	db.DB.Select(clause.Associations).Unscoped().Delete(&stand)
	return true
}
