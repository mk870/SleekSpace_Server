package services

import (
	"net/http"

	userDtos "SleekSpace/dtos/user"
	userModels "SleekSpace/models/user"
	userRepo "SleekSpace/repositories/user"
	"SleekSpace/utilities"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateUserProfilePicture(c *gin.Context) {
	var userProfilePicture userDtos.UserProfilePictureCreationDTO
	validateModelFields := validator.New()
	c.BindJSON(&userProfilePicture)

	modelFieldsValidationError := validateModelFields.Struct(userProfilePicture)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}
	user := userRepo.GetUserById(utilities.ConvertIntToString(userProfilePicture.UserId))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	newProfilePicture := userModels.UserProfilePicture{
		UserId:      userProfilePicture.UserId,
		Uri:         userProfilePicture.Uri,
		Name:        userProfilePicture.Name,
		FullPath:    userProfilePicture.FullPath,
		FileType:    userProfilePicture.FileType,
		ContentType: userProfilePicture.ContentType,
		Size:        userProfilePicture.Size,
	}
	isUserProfilePictureCreated := userRepo.CreateUserProfilePicture(user, &newProfilePicture)
	if isUserProfilePictureCreated {
		updatedUser := userRepo.GetUserById(utilities.ConvertIntToString(userProfilePicture.UserId))
		if updatedUser == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "this user account does not exist"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": utilities.UserResponseMapper(updatedUser, updatedUser.AccessToken)})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to create a profile picture"})
	}

}

func UpdateUserProfilePicture(c *gin.Context) {
	userId := c.Param("id")
	var profilePictureUpdate userDtos.UserProfilePictureResponseAndUpdateDTO
	validateModelFields := validator.New()
	c.BindJSON(&profilePictureUpdate)
	modelFieldsValidationError := validateModelFields.Struct(profilePictureUpdate)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}
	newProfilePicture := userModels.UserProfilePicture{
		Id:          profilePictureUpdate.Id,
		UserId:      profilePictureUpdate.UserId,
		Uri:         profilePictureUpdate.Uri,
		Name:        profilePictureUpdate.Name,
		FullPath:    profilePictureUpdate.FullPath,
		FileType:    profilePictureUpdate.FileType,
		ContentType: profilePictureUpdate.ContentType,
		Size:        profilePictureUpdate.Size,
	}
	isProfilePictureUpdated := userRepo.UpdateUserProfilePicture(&newProfilePicture)
	if !isProfilePictureUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "profile picture update failed"})
		return
	}
	updatedUser := userRepo.GetUserById(userId)
	if updatedUser == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user account does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": utilities.UserResponseMapper(updatedUser, updatedUser.AccessToken)})
}
