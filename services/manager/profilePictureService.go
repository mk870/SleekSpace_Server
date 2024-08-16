package services

import (
	"net/http"

	managerDtos "SleekSpace/dtos/manager"
	managerModels "SleekSpace/models/manager"
	managerRepo "SleekSpace/repositories/manager"
	"SleekSpace/utilities"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateManagerProfilePicture(c *gin.Context) {
	var managerProfilePicture managerDtos.ManagerProfilePictureCreationDTO
	validateModelFields := validator.New()
	c.BindJSON(&managerProfilePicture)

	modelFieldsValidationError := validateModelFields.Struct(managerProfilePicture)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}
	manager := managerRepo.GetManagerByManagerId(utilities.ConvertIntToString(managerProfilePicture.ManagerId))
	if manager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this management account does not exist"})
		return
	}
	newProfilePicture := managerModels.ManagerProfilePicture{
		ManagerId:   managerProfilePicture.ManagerId,
		Uri:         managerProfilePicture.Uri,
		Name:        managerProfilePicture.Name,
		FullPath:    managerProfilePicture.FullPath,
		FileType:    managerProfilePicture.FileType,
		ContentType: managerProfilePicture.ContentType,
		Size:        managerProfilePicture.Size,
	}
	isManagerProfilePictureCreated := managerRepo.CreateManagerProfilePicture(manager, &newProfilePicture)
	if isManagerProfilePictureCreated {
		updatedManager := managerRepo.GetManagerByManagerId(utilities.ConvertIntToString(managerProfilePicture.ManagerId))
		if updatedManager == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "this property management account does not exist"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": utilities.ManagerResponse(updatedManager)})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to create a profile picture"})
	}

}

func UpdateManagerProfilePicture(c *gin.Context) {
	managerId := c.Param("id")
	var profilePictureUpdate managerDtos.ManagerProfilePictureResponseAndUpdateDTO
	validateModelFields := validator.New()
	c.BindJSON(&profilePictureUpdate)
	modelFieldsValidationError := validateModelFields.Struct(profilePictureUpdate)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}
	newProfilePicture := managerModels.ManagerProfilePicture{
		Id:          profilePictureUpdate.Id,
		ManagerId:   profilePictureUpdate.ManagerId,
		Uri:         profilePictureUpdate.Uri,
		Name:        profilePictureUpdate.Name,
		FullPath:    profilePictureUpdate.FullPath,
		FileType:    profilePictureUpdate.FileType,
		ContentType: profilePictureUpdate.ContentType,
		Size:        profilePictureUpdate.Size,
	}
	isProfilePictureUpdated := managerRepo.UpdateProfilePicture(&newProfilePicture)
	if !isProfilePictureUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "profile picture update failed"})
		return
	}
	updatedManager := managerRepo.GetManagerByManagerId(managerId)
	if updatedManager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property management account does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": utilities.ManagerResponse(updatedManager)})
}
