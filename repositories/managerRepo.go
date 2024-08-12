package repositories

import (
	"SleekSpace/db"
	"SleekSpace/models"
	"errors"

	"gorm.io/gorm"
)

func CreateManager(user *models.User, manager *models.Manager) bool {
	err := db.DB.Model(user).Association("Manager").Append(manager)
	if err != nil {
		println(err.Error())
	}
	return true

}

func GetManagerByManagerId(managerId string) *models.Manager {
	var manager models.Manager
	result := db.DB.Preload("ManagerContactNumbers").First(&manager, managerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &manager
}

func UpdateManager(managerUpdate *models.Manager) bool {
	db.DB.Save(managerUpdate)
	return true
}

func GetManagerContactNumbersByManagerId(managerId int) []models.ManagerContactNumber {
	var manager = models.Manager{}
	result := db.DB.Preload("ManagerContactNumbers").First(&manager, managerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return manager.ManagerContactNumbers
}

func UpdateManagerContactNumbers(manager *models.Manager, updateManagerContactNumbersList []models.ManagerContactNumber) bool {
	manager.ManagerContactNumbers = updateManagerContactNumbersList
	db.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&manager)
	return true
}

func DeleteAllManagerContactNumbers(managerId int) bool {
	err := db.DB.Where("managerId = ?", managerId).Unscoped().Delete(&models.ManagerContactNumber{})
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return true
}

func DeleteManagerById(id string) bool {
	manager := GetManagerByManagerId(id)
	if manager == nil {
		return false
	}
	db.DB.Unscoped().Delete(&models.Manager{}, id)
	return true
}
