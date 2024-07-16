package services

import (
	"net/http"

	"SleekSpace/models"
	"SleekSpace/repositories"
	"SleekSpace/utilities"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	usersList := repositories.GetUsers()
	c.JSON(http.StatusOK, usersList)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user := repositories.GetUserById(id)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": utilities.UserResponseMapper(user, user.AccessToken)})
}

func GetUserByEmail(c *gin.Context) {
	client := c.MustGet("user").(*models.User)
	user := repositories.GetUserByEmail(client.Email)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": utilities.UserResponseMapper(user, user.AccessToken)})
}

func UpdateUser(c *gin.Context) {
	var update models.User
	c.BindJSON(&update)
	oldData := repositories.GetUserById(c.Param("id"))
	if oldData == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	oldData.Location = update.Location
	oldData.ContactNumber = update.ContactNumber
	oldData.WhatsAppNumber = update.WhatsAppNumber
	oldData.Avatar = update.Avatar
	updateResult := repositories.SaveUserUpdate(oldData)
	if !updateResult {
		c.JSON(http.StatusForbidden, gin.H{"error": "could not update the user data"})
		return
	}
	result := repositories.GetUserById(c.Param("id"))
	if result == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	if updateResult {
		c.JSON(http.StatusOK, gin.H{"response": utilities.UserResponseMapper(result, result.AccessToken)})
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	result := repositories.DeleteUserById(id)
	if result {
		c.String(http.StatusOK, "user successfully deleted")
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
	}
}
