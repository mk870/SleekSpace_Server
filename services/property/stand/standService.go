package stand

import (
	"net/http"

	managerDtos "SleekSpace/dtos/manager"
	managerModels "SleekSpace/models/manager"
	managerRepo "SleekSpace/repositories/manager"
	userRepo "SleekSpace/repositories/user"
	constantsUtilities "SleekSpace/utilities/constants"
	generalUtilities "SleekSpace/utilities/funcs/general"
	managerUtilities "SleekSpace/utilities/funcs/manager"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateStandForSale(c *gin.Context) {
	var manager managerDtos.ManagerCreationDTO
	validateModelFields := validator.New()
	c.BindJSON(&manager)

	modelFieldsValidationError := validateModelFields.Struct(manager)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	newManager := managerModels.Manager{
		UserId: manager.UserId,
		Name:   manager.Name,
		Email:  manager.Email,
		ProfilePicture: managerModels.ManagerProfilePicture{
			Name:        manager.ProfilePicture.Name,
			Uri:         manager.ProfilePicture.Uri,
			ContentType: manager.ProfilePicture.ContentType,
			FullPath:    manager.ProfilePicture.FullPath,
			FileType:    manager.ProfilePicture.FileType,
			Size:        manager.ProfilePicture.Size,
		},
		ManagerContactNumbers: []managerModels.ManagerContactNumber{
			{
				Number:       manager.Contacts[0].Number,
				CountryCode:  manager.Contacts[0].CountryCode,
				CountryAbbrv: manager.Contacts[0].CountryAbbrv,
				Type:         manager.Contacts[0].Type,
			},
			{
				Number:       manager.Contacts[1].Number,
				CountryCode:  manager.Contacts[1].CountryCode,
				CountryAbbrv: manager.Contacts[1].CountryAbbrv,
				Type:         manager.Contacts[1].Type,
			},
		},
	}
	isManagerCreated := managerRepo.CreateManager(&newManager)
	if isManagerCreated {
		createdManager := userRepo.GetUserByIdWithManager(generalUtilities.ConvertIntToString(manager.UserId))
		if createdManager == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.ManagerAccountCreatedButNoDataRetrievedError})
			return
		}
		createdManagerWithAssociations := managerRepo.GetManagerByManagerId(generalUtilities.ConvertIntToString(createdManager.Manager.Id))
		if createdManagerWithAssociations == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.ManagerAccountCreatedButNoDataRetrievedError})
		}
		c.JSON(http.StatusOK, gin.H{"response": managerUtilities.ManagerResponse(createdManagerWithAssociations)})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to create a manager"})
	}

}
