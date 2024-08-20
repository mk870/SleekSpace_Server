package commercial

import (
	"net/http"

	commercialDtos "SleekSpace/dtos/property/commercial"
	propertyModels "SleekSpace/models/property"
	commercialRepo "SleekSpace/repositories/property/commercial"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateCommercialRentalProperty(c *gin.Context) {
	var commercialRentalPropertyDetails commercialDtos.CommercialForRentPropertyCreationDto
	validateModelFields := validator.New()
	c.BindJSON(&commercialRentalPropertyDetails)

	modelFieldsValidationError := validateModelFields.Struct(commercialRentalPropertyDetails)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	newCommercialRentalProperty := propertyModels.CommercialRentalProperty{
		ManagerId:      commercialRentalPropertyDetails.ManagerId,
		UniqueId:       propertyUtilities.GeneratePropertyUniqueId(),
		RentAmount:     commercialRentalPropertyDetails.RentAmount,
		SizeNumber:     commercialRentalPropertyDetails.SizeNumber,
		SizeDimensions: commercialRentalPropertyDetails.SizeDimensions,
		Status:         commercialRentalPropertyDetails.Status,
		Type:           commercialRentalPropertyDetails.Type,
		YearBuilt:      commercialRentalPropertyDetails.YearBuilt,
		Stories:        commercialRentalPropertyDetails.Stories,
		HasElectricity: commercialRentalPropertyDetails.HasElectricity,
		HasWater:       commercialRentalPropertyDetails.HasWater,
		NumberOfRooms:  commercialRentalPropertyDetails.NumberOfRooms,
		IsFullSpace:    commercialRentalPropertyDetails.IsFullSpace,
		PropertyInsights: propertyModels.PropertyInsights{
			Views:             0,
			Shared:            0,
			AddedToFavourites: 0,
			ContactInfoViews:  0,
			PropertyType:      "commercialRentalProperty",
		},
		PropertyMedia: propertyUtilities.ConvertPropertyImagesOrVideosWithNoPropertyIdToModel(commercialRentalPropertyDetails.Media),
		Location: propertyModels.PropertyLocation{
			Boundingbox:  commercialRentalPropertyDetails.PropertyLocation.Boundingbox,
			Lat:          commercialRentalPropertyDetails.PropertyLocation.Lat,
			Lon:          commercialRentalPropertyDetails.PropertyLocation.Lon,
			DisplayName:  commercialRentalPropertyDetails.PropertyLocation.DisplayName,
			City:         commercialRentalPropertyDetails.PropertyLocation.City,
			County:       commercialRentalPropertyDetails.PropertyLocation.County,
			Country:      commercialRentalPropertyDetails.PropertyLocation.Country,
			CountryCode:  commercialRentalPropertyDetails.PropertyLocation.CountryCode,
			Province:     commercialRentalPropertyDetails.PropertyLocation.Province,
			Surburb:      commercialRentalPropertyDetails.PropertyLocation.Surburb,
			PropertyType: commercialRentalPropertyDetails.PropertyLocation.PropertyType,
		},
	}

	isCommercialRentalPropertyCreated := commercialRepo.CreateCommercialRentalProperty(&newCommercialRentalProperty)
	if isCommercialRentalPropertyCreated {
		c.JSON(http.StatusOK, gin.H{"response": "your property was successfully posted"})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to post your property"})
	}
}

func UpdateCommercialRentalPropertyDetails(c *gin.Context) {
	var commercialRentalPropertyUpdates commercialDtos.CommercialForRentPropertyUpdateDto
	validateModelFields := validator.New()
	c.BindJSON(&commercialRentalPropertyUpdates)

	modelFieldsValidationError := validateModelFields.Struct(commercialRentalPropertyUpdates)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	oldCommercialRentalPropertyData := commercialRepo.GetCommercialRentalPropertyById(c.Param("id"))
	if oldCommercialRentalPropertyData == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	}

	oldCommercialRentalPropertyData.RentAmount = commercialRentalPropertyUpdates.RentAmount
	oldCommercialRentalPropertyData.SizeNumber = commercialRentalPropertyUpdates.SizeNumber
	oldCommercialRentalPropertyData.SizeDimensions = commercialRentalPropertyUpdates.SizeDimensions
	oldCommercialRentalPropertyData.Status = commercialRentalPropertyUpdates.Status
	oldCommercialRentalPropertyData.Type = commercialRentalPropertyUpdates.Type
	oldCommercialRentalPropertyData.HasWater = commercialRentalPropertyUpdates.HasWater
	oldCommercialRentalPropertyData.HasElectricity = commercialRentalPropertyUpdates.HasElectricity
	oldCommercialRentalPropertyData.NumberOfRooms = commercialRentalPropertyUpdates.NumberOfRooms
	oldCommercialRentalPropertyData.Stories = commercialRentalPropertyUpdates.Stories
	oldCommercialRentalPropertyData.YearBuilt = commercialRentalPropertyUpdates.YearBuilt
	oldCommercialRentalPropertyData.IsFullSpace = commercialRentalPropertyUpdates.IsFullSpace
	oldCommercialRentalPropertyData.UniqueId = commercialRentalPropertyUpdates.UniqueId

	isCommercialRentalPropertyUpdated := commercialRepo.UpdateCommercialRentalProperty(oldCommercialRentalPropertyData)
	if !isCommercialRentalPropertyUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "property details update failed"})
		return
	}
	UpdateCommercialRentalProperty := commercialRepo.GetCommercialRentalPropertyWithAllAssociationsById(c.Param("id"))
	if UpdateCommercialRentalProperty == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.CommercialPropertyForRentResponse(*UpdateCommercialRentalProperty)})
}

func GetCommercialRentalPropertyId(c *gin.Context) {
	commercialRentalProperty := commercialRepo.GetCommercialRentalPropertyWithAllAssociationsById(c.Param("id"))
	if commercialRentalProperty == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.CommercialPropertyForRentResponse(*commercialRentalProperty)})
}

func GetManagerCommercialRentalPropertiesByManagerId(c *gin.Context) {
	properties := commercialRepo.GetManagerCommercialRentalPropertiesByManagerId(c.Param("id"))
	propertiesResponse := []commercialDtos.CommercialForRentPropertyResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			propertyResponse := propertyUtilities.CommercialPropertyForRentResponse(properties[i])
			propertiesResponse = append(propertiesResponse, propertyResponse)
		}
	}
	c.JSON(http.StatusOK, gin.H{"response": propertiesResponse})
}

func DeleteCommercialRentalPropertyById(c *gin.Context) {
	isPropertyDeleted := commercialRepo.DeleteCommercialRentalPropertyById(c.Param("id"))
	if !isPropertyDeleted {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"response": "your property was successfully deleted"})
		return
	}
}
