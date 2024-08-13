package repositories

import (
	"errors"

	"SleekSpace/db"
	"SleekSpace/models"

	"gorm.io/gorm"
	//"gorm.io/gorm/clause"
)

func CreateUser(user *models.User) bool {
	var existingUser models.User
	email := user.Email
	result := db.DB.Where("email =?", email).First(&existingUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		db.DB.Create(user)
		return true
	} else {
		return false
	}
}

func GetUsers() []models.User {
	var users = []models.User{}
	err := db.DB.Find(&users)
	if err != nil {
		println(err.Error, err.Name())
	}
	return users
}

func GetUserById(id string) *models.User {
	var user models.User
	result := db.DB.Preload("ContactNumbers").Preload("Location").First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func GetUserAndAllAssociationsById(id string) *models.User {
	var user models.User
	result := db.DB.Preload("ContactNumbers").Preload("Location").Preload("RegistrationCode").Preload("Manager").First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func GetUserByIdWithManager(id string) *models.User {
	var user models.User
	result := db.DB.Preload("Manager").First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func SaveUserUpdate(update *models.User) bool {
	db.DB.Save(update)
	return true
}

func DeleteUserById(id string) bool {
	user := GetUserById(id)
	if user == nil {
		return false
	}
	db.DB.Unscoped().Delete(&models.User{}, id)
	//db.DB.Select(clause.Associations).Unscoped().Delete(&models.User{}, id)
	return true
}

func GetUserByEmail(email string) *models.User {
	var user models.User
	result := db.DB.Where("email =?", email).Preload("ContactNumbers").Preload("Location").First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	} else {
		return &user
	}
}

func GetUserWithRegistrationCodeById(id string) models.User {
	var user models.User
	err := db.DB.Preload("RegistrationCode").First(&user, id)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user
}
