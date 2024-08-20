package stand

import (
	"net/http"

	standDtos "SleekSpace/dtos/property/stand"
	propertyModels "SleekSpace/models/property"
	standRepo "SleekSpace/repositories/property/stand"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateStandForSale(c *gin.Context) {
	var standInfo standDtos.StandCreationDTO
	validateModelFields := validator.New()
	c.BindJSON(&standInfo)

	modelFieldsValidationError := validateModelFields.Struct(standInfo)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	newStandForSale := propertyModels.PropertyStand{
		ManagerId:          standInfo.ManagerId,
		UniqueId:           propertyUtilities.GeneratePropertyUniqueId(),
		Price:              standInfo.Price,
		SizeNumber:         standInfo.SizeNumber,
		SizeDimensions:     standInfo.SizeDimensions,
		Status:             standInfo.Status,
		IsServiced:         standInfo.IsServiced,
		IsNegotiable:       standInfo.IsNegotiable,
		AreaHasElectricity: standInfo.AreaHasElectricity,
		Level:              standInfo.Level,
		Type:               standInfo.Type,
		PropertyInsights: propertyModels.PropertyInsights{
			Views:             0,
			Shared:            0,
			AddedToFavourites: 0,
			ContactInfoViews:  0,
			PropertyType:      "stand",
		},
		PropertyMedia: propertyUtilities.ConvertPropertyImagesOrVideosWithNoPropertyIdToModel(standInfo.Media),
		Location: propertyModels.PropertyLocation{
			Boundingbox:  standInfo.PropertyLocation.Boundingbox,
			Lat:          standInfo.PropertyLocation.Lat,
			Lon:          standInfo.PropertyLocation.Lon,
			DisplayName:  standInfo.PropertyLocation.DisplayName,
			City:         standInfo.PropertyLocation.City,
			County:       standInfo.PropertyLocation.County,
			Country:      standInfo.PropertyLocation.Country,
			CountryCode:  standInfo.PropertyLocation.CountryCode,
			Province:     standInfo.PropertyLocation.Province,
			Surburb:      standInfo.PropertyLocation.Surburb,
			PropertyType: standInfo.PropertyLocation.PropertyType,
		},
	}

	isStandCreated := standRepo.CreateStandForSale(&newStandForSale)
	if isStandCreated {
		c.JSON(http.StatusOK, gin.H{"response": "your stand was successfully posted"})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to post your stand"})
	}

}

func UpdateStandDetails(c *gin.Context) {
	var standUpdates standDtos.StandUpdateDTO
	validateModelFields := validator.New()
	c.BindJSON(&standUpdates)

	modelFieldsValidationError := validateModelFields.Struct(standUpdates)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	oldStandData := standRepo.GetStandById(c.Param("id"))
	if oldStandData == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this stand does not exist"})
		return
	}
	oldStandData.AreaHasElectricity = standUpdates.AreaHasElectricity
	oldStandData.IsNegotiable = standUpdates.IsNegotiable
	oldStandData.Price = standUpdates.Price
	oldStandData.SizeNumber = standUpdates.SizeNumber
	oldStandData.SizeDimensions = standUpdates.SizeDimensions
	oldStandData.Level = standUpdates.Level
	oldStandData.IsServiced = standUpdates.IsServiced
	oldStandData.Status = standUpdates.Status
	oldStandData.Type = standUpdates.Type
	oldStandData.UniqueId = standUpdates.UniqueId

	isStandUpdated := standRepo.UpdateStand(oldStandData)
	if !isStandUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "stand details update failed"})
		return
	}
	updatedStand := standRepo.GetStandWithAllAssociationsById(c.Param("id"))
	if updatedStand == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this stand does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.PropertyStandResponse(*updatedStand)})
}

func GetStandById(c *gin.Context) {
	stand := standRepo.GetStandWithAllAssociationsById(c.Param("id"))
	if stand == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this stand does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.PropertyStandResponse(*stand)})
}

func GetManagerStandsByManagerId(c *gin.Context) {
	stands := standRepo.GetManagerStandsByManagerId(c.Param("id"))
	standsResponse := []standDtos.StandResponseDTO{}
	if len(stands) > 0 {
		for i := 0; i < len(stands); i++ {
			standResponse := propertyUtilities.PropertyStandResponse(stands[i])
			standsResponse = append(standsResponse, standResponse)
		}
	}
	c.JSON(http.StatusOK, gin.H{"response": standsResponse})
}

func DeleteStandById(c *gin.Context) {
	isStandDeleted := standRepo.DeleteStandById(c.Param("id"))
	if !isStandDeleted {
		c.JSON(http.StatusForbidden, gin.H{"error": "this stand does not exist"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"response": "your stand was successfully deleted"})
		return
	}
}
