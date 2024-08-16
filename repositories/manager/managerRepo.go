package repositories

import (
	"SleekSpace/db"
	managerModels "SleekSpace/models/manager"
	userModels "SleekSpace/models/user"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateManager(user *userModels.User, manager *managerModels.Manager) bool {
	err := db.DB.Model(user).Association("Manager").Append(manager)
	if err != nil {
		println(err.Error())
	}
	return true

}

func GetManagerByManagerId(managerId string) *managerModels.Manager {
	var manager managerModels.Manager
	result := db.DB.Preload(clause.Associations).First(&manager, managerId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &manager
}

func UpdateManager(managerUpdate *managerModels.Manager) bool {
	db.DB.Save(managerUpdate)
	return true
}

func GetAllManagersContacts() []managerModels.ManagerContactNumber {
	var contacts = []managerModels.ManagerContactNumber{}
	err := db.DB.Find(&contacts)
	if err != nil {
		println(err.Error, err.Name())
	}
	return contacts
}

func GetAllManagers() []managerModels.Manager {
	var managers = []managerModels.Manager{}
	err := db.DB.Find(&managers)
	if err != nil {
		println(err.Error, err.Name())
	}
	return managers
}

func DeleteManagerById(id string) bool {
	manager := GetManagerByManagerId(id)
	if manager == nil {
		return false
	}
	db.DB.Select(clause.Associations).Unscoped().Delete(&manager)
	return true
}
