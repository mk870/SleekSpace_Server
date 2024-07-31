package services

import (
	"net/http"

	"SleekSpace/models"
	"SleekSpace/repositories"
	"SleekSpace/utilities"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateLocation(c *gin.Context) {
	var location models.Location
	validateModelFields := validator.New()
	c.BindJSON(&location)

	modelFieldsValidationError := validateModelFields.Struct(location)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}
	user := repositories.GetUserById(utilities.ConvertIntToString(location.UserId))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	isLocationCreated := repositories.CreateLocation(user, &location)
	if isLocationCreated {
		c.JSON(http.StatusOK, gin.H{"response": "location created succesfully"})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to create a location"})
	}

}

func UpdateLocation(c *gin.Context) {
	var location models.Location
	validateModelFields := validator.New()
	c.BindJSON(&location)
	modelFieldsValidationError := validateModelFields.Struct(location)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	isLocationUpdated := repositories.UpdateLocation(&location)
	if !isLocationUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "location update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "location was updated succesfully"})
}
