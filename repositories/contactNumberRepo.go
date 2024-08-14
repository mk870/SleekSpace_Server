package repositories

import (
	"SleekSpace/db"
	"SleekSpace/models"

	"gorm.io/gorm"
)

func CreateContactNumber(user *models.User, contactNumber *models.ContactNumber) bool {
	println("number: ", contactNumber.Number)
	err := db.DB.Model(user).Association("ContactNumbers").Append(contactNumber)
	if err != nil {
		println(err.Error())
	}
	return true
}

func GetAllUsersContactNumbers() []models.ContactNumber {
	var numbers = []models.ContactNumber{}
	err := db.DB.Find(&numbers)
	if err != nil {
		println(err.Error, err.Name())
	}
	return numbers
}

func GetUserContactNumbersByUserId(userId int) []models.ContactNumber {
	var user = models.User{}
	err := db.DB.Preload("ContactNumbers").First(&user, userId)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.ContactNumbers
}

func UpdateUserContactNumbers(user *models.User, updateContactNumbersList []models.ContactNumber) bool {
	user.ContactNumbers = updateContactNumbersList
	db.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
	return true
}
