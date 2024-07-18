package repositories

import (
	"SleekSpace/db"
	"SleekSpace/models"

	"gorm.io/gorm"
)

func CreateContactNumber(user *models.User, contactNumber models.ContactNumber) bool {
	err := db.DB.Model(user).Association("ContactNumber").Append(contactNumber)
	if err != nil {
		println(err.Error())
	}
	return true
}

func GetContactNumbers(id int) []models.ContactNumber {
	var user = models.User{}
	err := db.DB.Preload("ContactNumber").First(&user, id)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.ContactNumber
}

func UpdateProperty(user *models.User, updateContactNumbersList []models.ContactNumber) bool {
	user.ContactNumber = updateContactNumbersList
	db.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
	return true
}
