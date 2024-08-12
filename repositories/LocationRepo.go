package repositories

import (
	"SleekSpace/db"
	"SleekSpace/models"
)

func CreateLocation(user *models.User, location *models.Location) bool {
	err := db.DB.Model(user).Association("Location").Append(location)
	if err != nil {
		println(err.Error())
	}
	return true

}

func GetLocationById(userId string) models.Location {
	var user models.User
	err := db.DB.Preload("Location").First(&user, userId)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.Location
}

func UpdateLocation(locationUpdate *models.Location) bool {
	db.DB.Save(locationUpdate)
	return true
}

func DeleteLocation(userId int) bool {
	db.DB.Unscoped().Delete(&models.Location{}, userId)
	return true
}
