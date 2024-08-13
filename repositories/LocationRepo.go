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

func GetAllUsersLocations() []models.Location {
	var locations = []models.Location{}
	err := db.DB.Find(&locations)
	if err != nil {
		println(err.Error, err.Name())
	}
	return locations
}

func GetLocationByUserId(userId string) models.Location {
	var user models.User
	err := db.DB.Preload("Location").First(&user, userId)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.Location
}

func GetLocationById(locationId int) models.Location {
	var location = models.Location{}
	db.DB.First(&location, locationId)
	return location
}

func UpdateLocation(locationUpdate *models.Location) bool {
	db.DB.Save(locationUpdate)
	return true
}

func DeleteLocation(userId int) bool {
	db.DB.Unscoped().Delete(&models.Location{}, userId)
	return true
}
