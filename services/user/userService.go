package services

import (
	"net/http"

	userModels "SleekSpace/models/user"
	userRepo "SleekSpace/repositories/user"
	constantsUtilities "SleekSpace/utilities/constants"
	userUtilities "SleekSpace/utilities/funcs/user"

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
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": userUtilities.UserResponseMapper(user, user.AccessToken)})
}

func GetUserByEmail(c *gin.Context) {
	client := c.MustGet("user").(*userModels.User)
	user := userRepo.GetUserByEmail(client.Email)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": userUtilities.UserResponseMapper(user, user.AccessToken)})
}

func UpdateUser(c *gin.Context) {
	var update userModels.User
	c.BindJSON(&update)
	oldData := userRepo.GetUserById(c.Param("id"))
	if oldData == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	oldData.Location = update.Location
	oldData.ContactNumbers = update.ContactNumbers
	oldData.ProfilePicture = update.ProfilePicture
	updateResult := userRepo.SaveUserUpdate(oldData)
	if !updateResult {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.UserUpdateError})
		return
	}
	result := userRepo.GetUserById(c.Param("id"))
	if result == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	if updateResult {
		c.JSON(http.StatusOK, gin.H{"response": userUtilities.UserResponseMapper(result, result.AccessToken)})
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	isDeleted := userRepo.DeleteUserById(id)
	if isDeleted {
		c.String(http.StatusOK, constantsUtilities.UserDeletedSuccess)
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}

}
