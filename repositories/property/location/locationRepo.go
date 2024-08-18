package insights

import (
	"SleekSpace/db"
	propertyModels "SleekSpace/models/property"
)

func GetAllPropertyLocation() []propertyModels.PropertyLocation {
	var locations = []propertyModels.PropertyLocation{}
	err := db.DB.Find(&locations)
	if err != nil {
		println(err.Error, err.Name())
	}
	return locations
}

func GetPropertyLocationById(propertyLocationId int) propertyModels.PropertyLocation {
	var location = propertyModels.PropertyLocation{}
	db.DB.First(&location, propertyLocationId)
	return location
}

func UpdatePropertyLocation(propertyLocationUpdate *propertyModels.PropertyLocation) bool {
	db.DB.Save(propertyLocationUpdate)
	return true
}

func DeletePropertyLocation(propertyLocationId int) bool {
	db.DB.Unscoped().Delete(&propertyModels.PropertyLocation{}, propertyLocationId)
	return true
}
