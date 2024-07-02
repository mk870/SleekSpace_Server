package services

import (
	"net/http"

	"SleekSpace/models"
	"SleekSpace/repositories"

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

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var update models.User
	c.BindJSON(&update)
	oldData := repositories.GetUserById(c.Param("id"))
	if oldData == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	if update.GivenName != "" {
		oldData.GivenName = update.GivenName
	}
	if update.FamilyName != "" {
		oldData.FamilyName = update.FamilyName
	}
	updateResult := repositories.SaveUserUpdate(oldData)
	if updateResult {
		c.String(http.StatusOK, "user data updated successfully")
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
