package services

import (
	"net/http"

	userModels "SleekSpace/models/user"
	userRepo "SleekSpace/repositories/user"
	"SleekSpace/utilities"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	usersList := userRepo.GetUsers()
	c.JSON(http.StatusOK, usersList)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user := userRepo.GetUserById(id)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": utilities.NoUserError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": utilities.UserResponseMapper(user, user.AccessToken)})
}

func GetUserByEmail(c *gin.Context) {
	client := c.MustGet("user").(*userModels.User)
	user := userRepo.GetUserByEmail(client.Email)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": utilities.NoUserError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": utilities.UserResponseMapper(user, user.AccessToken)})
}

func UpdateUser(c *gin.Context) {
	var update userModels.User
	c.BindJSON(&update)
	oldData := userRepo.GetUserById(c.Param("id"))
	if oldData == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": utilities.NoUserError})
		return
	}
	oldData.Location = update.Location
	oldData.ContactNumbers = update.ContactNumbers
	oldData.ProfilePicture = update.ProfilePicture
	updateResult := userRepo.SaveUserUpdate(oldData)
	if !updateResult {
		c.JSON(http.StatusForbidden, gin.H{"error": utilities.UserUpdateError})
		return
	}
	result := userRepo.GetUserById(c.Param("id"))
	if result == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": utilities.NoUserError})
		return
	}
	if updateResult {
		c.JSON(http.StatusOK, gin.H{"response": utilities.UserResponseMapper(result, result.AccessToken)})
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	isDeleted := userRepo.DeleteUserById(id)
	if isDeleted {
		c.String(http.StatusOK, utilities.UserDeletedSuccess)
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": utilities.NoUserError})
		return
	}

}