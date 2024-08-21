package land

import (
	"net/http"

	landDtos "SleekSpace/dtos/property/land"
	propertyModels "SleekSpace/models/property"
	landRepo "SleekSpace/repositories/property/land"
	constants "SleekSpace/utilities/constants"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateLandPropertyForSale(c *gin.Context) {
	var landDetails landDtos.LandForSalePropertyCreationDto
	validateModelFields := validator.New()
	c.BindJSON(&landDetails)

	modelFieldsValidationError := validateModelFields.Struct(landDetails)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	newLandForSale := propertyModels.LandForSaleProperty{
		ManagerId:          landDetails.ManagerId,
		UniqueId:           propertyUtilities.GeneratePropertyUniqueId(),
		Price:              landDetails.Price,
		SizeNumber:         landDetails.SizeNumber,
		SizeDimensions:     landDetails.SizeDimensions,
		Status:             landDetails.Status,
		Type:               landDetails.Type,
		AreaHasElectricity: landDetails.AreaHasElectricity,
		HasWater:           landDetails.HasWater,
		IsNegotiable:       landDetails.IsNegotiable,
		PropertyInsights: propertyModels.PropertyInsights{
			Views:             0,
			Shared:            0,
			AddedToFavourites: 0,
			ContactInfoViews:  0,
			PropertyType:      constants.LandPropertyType,
		},
		PropertyMedia: propertyUtilities.ConvertPropertyImagesOrVideosWithNoPropertyIdToModel(landDetails.Media, constants.LandPropertyType),
		Location: propertyModels.PropertyLocation{
			Boundingbox:  landDetails.PropertyLocation.Boundingbox,
			Lat:          landDetails.PropertyLocation.Lat,
			Lon:          landDetails.PropertyLocation.Lon,
			DisplayName:  landDetails.PropertyLocation.DisplayName,
			City:         landDetails.PropertyLocation.City,
			County:       landDetails.PropertyLocation.County,
			Country:      landDetails.PropertyLocation.Country,
			CountryCode:  landDetails.PropertyLocation.CountryCode,
			Province:     landDetails.PropertyLocation.Province,
			Surburb:      landDetails.PropertyLocation.Surburb,
			PropertyType: constants.LandPropertyType,
		},
	}

	isLandCreated := landRepo.CreateLandPropertyForSale(&newLandForSale)
	if isLandCreated {
		c.JSON(http.StatusOK, gin.H{"response": "your land was successfully posted"})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to post your land"})
	}

}

func UpdateLandPropertyDetails(c *gin.Context) {
	var landUpdates landDtos.LandForSalePropertyUpdateDto
	validateModelFields := validator.New()
	c.BindJSON(&landUpdates)

	modelFieldsValidationError := validateModelFields.Struct(landUpdates)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	oldLandData := landRepo.GetLandPropertyForSaleById(c.Param("id"))
	if oldLandData == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this land does not exist"})
		return
	}
	oldLandData.AreaHasElectricity = landUpdates.AreaHasElectricity
	oldLandData.HasWater = landUpdates.HasWater
	oldLandData.Price = landUpdates.Price
	oldLandData.SizeNumber = landUpdates.SizeNumber
	oldLandData.SizeDimensions = landUpdates.SizeDimensions
	oldLandData.Status = landUpdates.Status
	oldLandData.Type = landUpdates.Type
	oldLandData.UniqueId = landUpdates.UniqueId
	oldLandData.IsNegotiable = landUpdates.IsNegotiable

	isStandUpdated := landRepo.UpdateLandPropertyForSale(oldLandData)
	if !isStandUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "land details update failed"})
		return
	}
	updatedLand := landRepo.GetLandPropertyForSaleWithAllAssociationsById(c.Param("id"))
	if updatedLand == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this land does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.LandPropertyResponse(*updatedLand)})
}

func GetLandPropertyById(c *gin.Context) {
	stand := landRepo.GetLandPropertyForSaleWithAllAssociationsById(c.Param("id"))
	if stand == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this land does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.LandPropertyResponse(*stand)})
}

func GetManagerLandPropertiesByManagerId(c *gin.Context) {
	landProperties := landRepo.GetManagerLandPropertiesForSaleByManagerId(c.Param("id"))
	landPropertiesResponse := []landDtos.LandForSalePropertyResponseDto{}
	if len(landProperties) > 0 {
		for i := 0; i < len(landProperties); i++ {
			landResponse := propertyUtilities.LandPropertyResponse(landProperties[i])
			landPropertiesResponse = append(landPropertiesResponse, landResponse)
		}
	}
	c.JSON(http.StatusOK, gin.H{"response": landPropertiesResponse})
}

func DeleteLandPropertyById(c *gin.Context) {
	isLandDeleted := landRepo.DeleteLandPropertyForSaleById(c.Param("id"))
	if !isLandDeleted {
		c.JSON(http.StatusForbidden, gin.H{"error": "this land does not exist"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"response": "your land was successfully deleted"})
		return
	}
}
