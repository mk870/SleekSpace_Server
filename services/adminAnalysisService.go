package services

import (
	"net/http"

	"SleekSpace/repositories"
	"SleekSpace/utilities"

	"github.com/gin-gonic/gin"
)

func GetLocationById(c *gin.Context) {
	id := c.Param("id")
	location := repositories.GetLocationById(utilities.ConvertStringToInt(id))
	c.JSON(http.StatusOK, gin.H{"response": location})
}

func GetAllUsersLocations(c *gin.Context) {
	codes := repositories.GetAllUsersLocations()
	c.JSON(http.StatusOK, gin.H{
		"response": codes,
	})
}

func GetVerificationCodeById(c *gin.Context) {
	id := c.Param("id")
	code := repositories.GetVerificationCodeById(id)
	c.JSON(http.StatusOK, gin.H{
		"response": code,
	})
}

func GetAllVerificationCodes(c *gin.Context) {
	codes := repositories.AllVerificationCodes()
	c.JSON(http.StatusOK, gin.H{
		"response": codes,
	})
}

func DeleteVerificationCode(c *gin.Context) {
	id := c.Param("id")
	isDeleted := repositories.DeleteVerficationCode(utilities.ConvertStringToInt(id))
	if isDeleted {
		c.JSON(http.StatusOK, gin.H{
			"response": "code deleted",
		})
	}
}

func GetUserContacts(c *gin.Context) {
	numbers := repositories.GetAllUsersContactNumbers()
	c.JSON(http.StatusOK, gin.H{
		"response": numbers,
	})
}

func GetAllManagersContacts(c *gin.Context) {
	numbers := repositories.GetAllManagersContacts()
	c.JSON(http.StatusOK, gin.H{
		"response": numbers,
	})
}

func GetAllManagers(c *gin.Context) {
	managers := repositories.GetAllManagers()
	c.JSON(http.StatusOK, gin.H{
		"response": managers,
	})
}
