package services

import (
	"net/http"

	"SleekSpace/models"
	"SleekSpace/repositories"
	"SleekSpace/utilities"

	"github.com/gin-gonic/gin"
)

func CreateContactNumber(c *gin.Context) {
	var contact models.ContactNumber
	c.BindJSON(&contact)
	user := repositories.GetUserById(utilities.ConvertIntToString(contact.UserId))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	result := repositories.CreateContactNumber(user, &contact)
	if result {
		c.String(http.StatusOK, "contact added")
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "an error occured"})
	}
}

func UpdateContactNumbers(c *gin.Context) {
	type ContactNumbers struct {
		Contacts []models.ContactNumber `json:"contacts"`
	}
	var contacts ContactNumbers
	c.BindJSON(&contacts)
	user := repositories.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	isContactsUpdated := repositories.UpdateUserContactNumbers(user, contacts.Contacts)
	if !isContactsUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "contacts update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "contacts were updated succesfully"})
}
