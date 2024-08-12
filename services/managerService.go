package services

import (
	"net/http"

	"SleekSpace/dtos"
	"SleekSpace/models"
	"SleekSpace/repositories"
	"SleekSpace/utilities"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateManager(c *gin.Context) {
	var manager dtos.ManagerDTO
	validateModelFields := validator.New()
	c.BindJSON(&manager)

	modelFieldsValidationError := validateModelFields.Struct(manager)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}
	user := repositories.GetUserById(utilities.ConvertIntToString(manager.UserId))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	newManager := models.Manager{
		UserId: manager.UserId,
		Name:   manager.Name,
		Email:  manager.Email,
		Avatar: manager.Avatar,
	}
	isManagerCreated := repositories.CreateManager(user, &newManager)
	if isManagerCreated {
		userWithManager := repositories.GetUserByIdWithManager(utilities.ConvertIntToString(user.Id))
		if userWithManager == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
			return
		}
		managerContacts := utilities.AddManagerIdToContacts(manager.Contacts, userWithManager.Manager.Id)
		isManagerContactNumberCreated := repositories.UpdateManagerContactNumbers(&userWithManager.Manager, managerContacts)
		if !isManagerContactNumberCreated {
			c.JSON(http.StatusForbidden, gin.H{"error": "failed to create a manager contact number"})
			return
		}
		createdmanager := repositories.GetManagerByManagerId(utilities.ConvertIntToString(userWithManager.Manager.Id))
		c.JSON(http.StatusOK, gin.H{"response": utilities.ManagerResponse(createdmanager)})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to create a manager"})
	}

}

func GetManagerByUserId(c *gin.Context) {
	id := c.Param("userId")
	user := repositories.GetUserByIdWithManager(id)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	manager := repositories.GetManagerByManagerId(utilities.ConvertIntToString(user.Manager.Id))
	if manager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property management account does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": utilities.ManagerResponse(manager)})
}

func GetManagerByManagerId(c *gin.Context) {
	id := c.Param("id")
	manager := repositories.GetManagerByManagerId(id)
	if manager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property management account does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": utilities.ManagerResponse(manager)})
}

func UpdateManager(c *gin.Context) {
	type ManagerUpdate struct {
		Id       int                           `json:"id"`
		UserId   int                           `json:"userId"`
		Name     string                        `json:"name" validate:"required,min=2,max=50"`
		Email    string                        `json:"email"`
		Avatar   string                        `json:"avatar"`
		Contacts []models.ManagerContactNumber `json:"contacts"`
	}
	var managerUpdate ManagerUpdate
	c.BindJSON(&managerUpdate)
	manager := repositories.GetManagerByManagerId(c.Param("id"))
	if manager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property management account does not exist"})
		return
	}
	isContactsUpdated := repositories.UpdateManagerContactNumbers(manager, managerUpdate.Contacts)
	if !isContactsUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "property management account contacts failed to update"})
		return
	}
	updatedManagerContacts := repositories.GetManagerContactNumbersByManagerId(manager.Id)
	if updatedManagerContacts == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "property management account contact numbers do not exist"})
		return
	}
	manager.Avatar = managerUpdate.Avatar
	manager.Email = managerUpdate.Email
	manager.Name = managerUpdate.Name
	manager.ManagerContactNumbers = updatedManagerContacts
	isManagerUpdated := repositories.UpdateManager(manager)
	if isManagerUpdated {
		updatedManager := repositories.GetManagerByManagerId(c.Param("id"))
		if updatedManager == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "this property management account does not exist"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": utilities.ManagerResponse(updatedManager)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": " this property management account failed to update"})
}

func DeleteAllManagerContactNumbers(c *gin.Context) {
	id := c.Param("id")
	result := repositories.DeleteAllManagerContactNumbers(utilities.ConvertStringToInt(id))
	if result {
		c.String(http.StatusOK, "your property management account was successfully deleted")
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property management account does not exist"})
	}
}

func DeleteManager(c *gin.Context) {
	id := c.Param("id")
	areManagerContactsDeleted := repositories.DeleteAllManagerContactNumbers(utilities.ConvertStringToInt(id))
	if areManagerContactsDeleted {
		result := repositories.DeleteManagerById(id)
		if result {
			c.String(http.StatusOK, "your property management account was successfully deleted")
			return
		} else {
			c.JSON(http.StatusForbidden, gin.H{"error": "this property management account does not exist"})
			return
		}
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to delete manager account details, please try again later"})
		return
	}
}
