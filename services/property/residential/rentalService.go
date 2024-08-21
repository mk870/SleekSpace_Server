package residential

import (
	"net/http"

	residentialDtos "SleekSpace/dtos/property/residential"
	propertyModels "SleekSpace/models/property"
	residentialRepo "SleekSpace/repositories/property/residential"
	constants "SleekSpace/utilities/constants"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateResidentialRentalProperty(c *gin.Context) {
	var residentialRentalPropertyDetails residentialDtos.ResidentialPropertyForRentCreationDto
	validateModelFields := validator.New()
	c.BindJSON(&residentialRentalPropertyDetails)

	modelFieldsValidationError := validateModelFields.Struct(residentialRentalPropertyDetails)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	newResidentialRentalProperty := propertyModels.ResidentialRentalProperty{
		ManagerId:       residentialRentalPropertyDetails.ManagerId,
		UniqueId:        propertyUtilities.GeneratePropertyUniqueId(),
		RentAmount:      residentialRentalPropertyDetails.RentAmount,
		SizeNumber:      residentialRentalPropertyDetails.SizeNumber,
		SizeDimensions:  residentialRentalPropertyDetails.SizeDimensions,
		Status:          residentialRentalPropertyDetails.Status,
		Type:            residentialRentalPropertyDetails.Type,
		YearBuilt:       residentialRentalPropertyDetails.YearBuilt,
		Bedrooms:        residentialRentalPropertyDetails.Bedrooms,
		Bathrooms:       residentialRentalPropertyDetails.Bathrooms,
		Stories:         residentialRentalPropertyDetails.Stories,
		HasElectricity:  residentialRentalPropertyDetails.HasElectricity,
		HasWater:        residentialRentalPropertyDetails.HasWater,
		NumberOfRooms:   residentialRentalPropertyDetails.NumberOfRooms,
		NumberOfGarages: residentialRentalPropertyDetails.NumberOfGarages,
		HasSwimmingPool: residentialRentalPropertyDetails.HasSwimmingPool,
		IsFullHouse:     residentialRentalPropertyDetails.IsFullHouse,
		PropertyInsights: propertyModels.PropertyInsights{
			Views:             0,
			Shared:            0,
			AddedToFavourites: 0,
			ContactInfoViews:  0,
			PropertyType:      constants.ResidentialRentalPropertyType,
		},
		PropertyMedia: propertyUtilities.ConvertPropertyImagesOrVideosWithNoPropertyIdToModel(residentialRentalPropertyDetails.Media, constants.ResidentialRentalPropertyType),
		Location: propertyModels.PropertyLocation{
			Boundingbox:  residentialRentalPropertyDetails.PropertyLocation.Boundingbox,
			Lat:          residentialRentalPropertyDetails.PropertyLocation.Lat,
			Lon:          residentialRentalPropertyDetails.PropertyLocation.Lon,
			DisplayName:  residentialRentalPropertyDetails.PropertyLocation.DisplayName,
			City:         residentialRentalPropertyDetails.PropertyLocation.City,
			County:       residentialRentalPropertyDetails.PropertyLocation.County,
			Country:      residentialRentalPropertyDetails.PropertyLocation.Country,
			CountryCode:  residentialRentalPropertyDetails.PropertyLocation.CountryCode,
			Province:     residentialRentalPropertyDetails.PropertyLocation.Province,
			Surburb:      residentialRentalPropertyDetails.PropertyLocation.Surburb,
			PropertyType: constants.ResidentialRentalPropertyType,
		},
	}

	isResidentialRentalPropertyCreated := residentialRepo.CreateResidentialRentalProperty(&newResidentialRentalProperty)
	if isResidentialRentalPropertyCreated {
		c.JSON(http.StatusOK, gin.H{"response": "your property was successfully posted"})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to post your property"})
	}
}

func UpdateResidentialRentalPropertyDetails(c *gin.Context) {
	var residentialRentalPropertyUpdates residentialDtos.ResidentialPropertyForRentUpdateDto
	validateModelFields := validator.New()
	c.BindJSON(&residentialRentalPropertyUpdates)

	modelFieldsValidationError := validateModelFields.Struct(residentialRentalPropertyUpdates)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	oldResidentialRentalPropertyData := residentialRepo.GetResidentialRentalPropertyById(c.Param("id"))
	if oldResidentialRentalPropertyData == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	}

	oldResidentialRentalPropertyData.IsFullHouse = residentialRentalPropertyUpdates.IsFullHouse
	oldResidentialRentalPropertyData.RentAmount = residentialRentalPropertyUpdates.RentAmount
	oldResidentialRentalPropertyData.SizeNumber = residentialRentalPropertyUpdates.SizeNumber
	oldResidentialRentalPropertyData.SizeDimensions = residentialRentalPropertyUpdates.SizeDimensions
	oldResidentialRentalPropertyData.Status = residentialRentalPropertyUpdates.Status
	oldResidentialRentalPropertyData.Type = residentialRentalPropertyUpdates.Type
	oldResidentialRentalPropertyData.Bathrooms = residentialRentalPropertyUpdates.Bathrooms
	oldResidentialRentalPropertyData.Bedrooms = residentialRentalPropertyUpdates.Bedrooms
	oldResidentialRentalPropertyData.HasSwimmingPool = residentialRentalPropertyUpdates.HasSwimmingPool
	oldResidentialRentalPropertyData.HasWater = residentialRentalPropertyUpdates.HasWater
	oldResidentialRentalPropertyData.HasElectricity = residentialRentalPropertyUpdates.HasElectricity
	oldResidentialRentalPropertyData.NumberOfRooms = residentialRentalPropertyUpdates.NumberOfRooms
	oldResidentialRentalPropertyData.NumberOfGarages = residentialRentalPropertyUpdates.NumberOfGarages
	oldResidentialRentalPropertyData.Stories = residentialRentalPropertyUpdates.Stories
	oldResidentialRentalPropertyData.YearBuilt = residentialRentalPropertyUpdates.YearBuilt
	oldResidentialRentalPropertyData.UniqueId = residentialRentalPropertyUpdates.UniqueId

	isResidentialRentalPropertyUpdated := residentialRepo.UpdateResidentialRentalProperty(oldResidentialRentalPropertyData)
	if !isResidentialRentalPropertyUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "property details update failed"})
		return
	}
	UpdateResidentialRentalProperty := residentialRepo.GetResidentialRentalPropertyWithAllAssociationsById(c.Param("id"))
	if UpdateResidentialRentalProperty == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.ResidentialRentalPropertyResponse(*UpdateResidentialRentalProperty)})
}

func GetResidentialRentalPropertyId(c *gin.Context) {
	residentialRentalProperty := residentialRepo.GetResidentialRentalPropertyWithAllAssociationsById(c.Param("id"))
	if residentialRentalProperty == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.ResidentialRentalPropertyResponse(*residentialRentalProperty)})
}

func GetManagerResidentialRentalPropertiesByManagerId(c *gin.Context) {
	properties := residentialRepo.GetManagerResidentialRentalPropertiesByManagerId(c.Param("id"))
	propertiesResponse := []residentialDtos.ResidentialPropertyForRentResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			propertyResponse := propertyUtilities.ResidentialRentalPropertyResponse(properties[i])
			propertiesResponse = append(propertiesResponse, propertyResponse)
		}
	}
	c.JSON(http.StatusOK, gin.H{"response": propertiesResponse})
}

func DeleteResidentialRentalPropertyById(c *gin.Context) {
	isPropertyDeleted := residentialRepo.DeleteResidentialRentalPropertyById(c.Param("id"))
	if !isPropertyDeleted {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"response": "your property was successfully deleted"})
		return
	}
}
