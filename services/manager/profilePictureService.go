package services

import (
	"net/http"

	managerDtos "SleekSpace/dtos/manager"
	managerModels "SleekSpace/models/manager"
	managerRepo "SleekSpace/repositories/manager"
	constantsUtilities "SleekSpace/utilities/constants"
	managerUtilities "SleekSpace/utilities/funcs/manager"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

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
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoManagerAccountError})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": managerUtilities.ManagerResponse(updatedManager)})
}
