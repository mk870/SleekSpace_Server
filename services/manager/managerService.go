package services

import (
	"net/http"

	managerDtos "SleekSpace/dtos/manager"
	managerModels "SleekSpace/models/manager"
	managerRepo "SleekSpace/repositories/manager"
	userRepo "SleekSpace/repositories/user"
	"SleekSpace/utilities"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateManager(c *gin.Context) {
	var manager managerDtos.ManagerCreationDTO
	validateModelFields := validator.New()
	c.BindJSON(&manager)

	modelFieldsValidationError := validateModelFields.Struct(manager)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}
	user := userRepo.GetUserById(utilities.ConvertIntToString(manager.UserId))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	newManager := managerModels.Manager{
		UserId: manager.UserId,
		Name:   manager.Name,
		Email:  manager.Email,
	}
	isManagerCreated := managerRepo.CreateManager(user, &newManager)
	if isManagerCreated {
		createdManager := userRepo.GetUserByIdWithManager(utilities.ConvertIntToString(manager.UserId))
		c.JSON(http.StatusOK, gin.H{"response": utilities.ManagerResponse(&createdManager.Manager)})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to create a manager"})
	}

}

func GetManagerByUserId(c *gin.Context) {
	id := c.Param("userId")
	user := userRepo.GetUserByIdWithManager(id)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	manager := managerRepo.GetManagerByManagerId(utilities.ConvertIntToString(user.Manager.Id))
	if manager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property management account does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": utilities.ManagerResponse(manager)})
}

func GetManagerByManagerId(c *gin.Context) {
	id := c.Param("id")
	manager := managerRepo.GetManagerByManagerId(id)
	if manager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property management account does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": utilities.ManagerResponse(manager)})
}

func UpdateManagerEmailAndName(c *gin.Context) {
	managerId := c.Param("id")
	var nameAndEmailUpdates managerDtos.UpdateManagerNameAndEmailDTO
	validateModelFields := validator.New()
	c.BindJSON(&nameAndEmailUpdates)

	modelFieldsValidationError := validateModelFields.Struct(nameAndEmailUpdates)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	manager := managerRepo.GetManagerByManagerId(managerId)
	if manager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property management account does not exist"})
		return
	}

	manager.Email = nameAndEmailUpdates.Email
	manager.Name = nameAndEmailUpdates.Name
	isManagerUpdated := managerRepo.UpdateManager(manager)
	if isManagerUpdated {
		updatedManager := managerRepo.GetManagerByManagerId(c.Param("id"))
		if updatedManager == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "this property management account does not exist"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": utilities.ManagerResponse(updatedManager)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": " this property management account failed to update"})

}

func DeleteManager(c *gin.Context) {
	id := c.Param("id")
	isManagerDeleted := managerRepo.DeleteManagerById(id)
	if isManagerDeleted {
		c.String(http.StatusOK, "your property management account was successfully deleted")
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property management account does not exist"})
		return
	}
}
