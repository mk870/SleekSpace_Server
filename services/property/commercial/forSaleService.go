package commercial

import (
	"net/http"

	commercialDtos "SleekSpace/dtos/property/commercial"
	managerModels "SleekSpace/models/manager"
	propertyModels "SleekSpace/models/property"
	managerRepo "SleekSpace/repositories/manager"
	commercialRepo "SleekSpace/repositories/property/commercial"
	constants "SleekSpace/utilities/constants"
	generalUtilities "SleekSpace/utilities/funcs/general"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateCommercialPropertyForSale(c *gin.Context) {
	var commercialPropertyForSaleDetails commercialDtos.CommercialForSalePropertyCreationDto
	validateModelFields := validator.New()
	c.BindJSON(&commercialPropertyForSaleDetails)

	modelFieldsValidationError := validateModelFields.Struct(commercialPropertyForSaleDetails)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	manager := managerRepo.GetManagerWithProfilePictureAndContactsByManagerId(generalUtilities.ConvertIntToString(commercialPropertyForSaleDetails.ManagerId))
	if manager == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this manager does not exist"})
		return
	}

	newCommercialPropertyForSale := managerModels.CommercialForSaleProperty{
		ManagerId:        commercialPropertyForSaleDetails.ManagerId,
		UniqueId:         propertyUtilities.GeneratePropertyUniqueId(),
		Price:            commercialPropertyForSaleDetails.Price,
		SizeNumber:       commercialPropertyForSaleDetails.SizeNumber,
		SizeDimensions:   commercialPropertyForSaleDetails.SizeDimensions,
		Status:           commercialPropertyForSaleDetails.Status,
		Type:             commercialPropertyForSaleDetails.Type,
		YearBuilt:        commercialPropertyForSaleDetails.YearBuilt,
		Stories:          commercialPropertyForSaleDetails.Stories,
		HasElectricity:   commercialPropertyForSaleDetails.HasElectricity,
		HasWater:         commercialPropertyForSaleDetails.HasWater,
		IsNegotiable:     commercialPropertyForSaleDetails.IsNegotiable,
		NumberOfRooms:    commercialPropertyForSaleDetails.NumberOfRooms,
		ExteriorFeatures: commercialPropertyForSaleDetails.ExteriorFeatures,
		InteriorFeatures: commercialPropertyForSaleDetails.InteriorFeatures,
		OtherDetails:     commercialPropertyForSaleDetails.OtherDetails,
		Manager:          *manager,
		PropertyInsights: propertyModels.PropertyInsights{
			Views:             0,
			Shared:            0,
			AddedToFavourites: 0,
			ContactInfoViews:  0,
			PropertyType:      constants.CommercialPropertyForSaleType,
		},
		PropertyMedia: propertyUtilities.ConvertPropertyImagesOrVideosWithNoPropertyIdToModel(commercialPropertyForSaleDetails.Media, constants.CommercialPropertyForSaleType),
		Location: propertyModels.PropertyLocation{
			Boundingbox:  commercialPropertyForSaleDetails.PropertyLocation.Boundingbox,
			Lat:          commercialPropertyForSaleDetails.PropertyLocation.Lat,
			Lon:          commercialPropertyForSaleDetails.PropertyLocation.Lon,
			DisplayName:  commercialPropertyForSaleDetails.PropertyLocation.DisplayName,
			City:         commercialPropertyForSaleDetails.PropertyLocation.City,
			County:       commercialPropertyForSaleDetails.PropertyLocation.County,
			Country:      commercialPropertyForSaleDetails.PropertyLocation.Country,
			CountryCode:  commercialPropertyForSaleDetails.PropertyLocation.CountryCode,
			Province:     commercialPropertyForSaleDetails.PropertyLocation.Province,
			Surburb:      commercialPropertyForSaleDetails.PropertyLocation.Surburb,
			PropertyType: constants.CommercialPropertyForSaleType,
		},
	}

	isCommercialPropertyForSaleCreated := commercialRepo.CreateCommercialPropertyForSale(&newCommercialPropertyForSale)
	if isCommercialPropertyForSaleCreated {
		c.JSON(http.StatusOK, gin.H{"response": "your property was successfully posted"})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to post your property"})
	}
}

func UpdateCommercialPropertyForSaleDetails(c *gin.Context) {
	var commercialPropertyForSaleUpdates commercialDtos.CommercialForSalePropertyUpdateDto
	validateModelFields := validator.New()
	c.BindJSON(&commercialPropertyForSaleUpdates)

	modelFieldsValidationError := validateModelFields.Struct(commercialPropertyForSaleUpdates)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	oldCommercialPropertyForSaleData := commercialRepo.GetCommercialPropertyForSaleById(c.Param("id"))
	if oldCommercialPropertyForSaleData == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	}

	oldCommercialPropertyForSaleData.Price = commercialPropertyForSaleUpdates.Price
	oldCommercialPropertyForSaleData.SizeNumber = commercialPropertyForSaleUpdates.SizeNumber
	oldCommercialPropertyForSaleData.SizeDimensions = commercialPropertyForSaleUpdates.SizeDimensions
	oldCommercialPropertyForSaleData.Status = commercialPropertyForSaleUpdates.Status
	oldCommercialPropertyForSaleData.Type = commercialPropertyForSaleUpdates.Type
	oldCommercialPropertyForSaleData.HasWater = commercialPropertyForSaleUpdates.HasWater
	oldCommercialPropertyForSaleData.HasElectricity = commercialPropertyForSaleUpdates.HasElectricity
	oldCommercialPropertyForSaleData.NumberOfRooms = commercialPropertyForSaleUpdates.NumberOfRooms
	oldCommercialPropertyForSaleData.Stories = commercialPropertyForSaleUpdates.Stories
	oldCommercialPropertyForSaleData.YearBuilt = commercialPropertyForSaleUpdates.YearBuilt
	oldCommercialPropertyForSaleData.UniqueId = commercialPropertyForSaleUpdates.UniqueId
	oldCommercialPropertyForSaleData.IsNegotiable = commercialPropertyForSaleUpdates.IsNegotiable
	oldCommercialPropertyForSaleData.InteriorFeatures = commercialPropertyForSaleUpdates.InteriorFeatures
	oldCommercialPropertyForSaleData.ExteriorFeatures = commercialPropertyForSaleUpdates.ExteriorFeatures
	oldCommercialPropertyForSaleData.OtherDetails = commercialPropertyForSaleUpdates.OtherDetails

	isCommercialPropertyForSaleUpdated := commercialRepo.UpdateCommercialPropertyForSale(oldCommercialPropertyForSaleData)
	if !isCommercialPropertyForSaleUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "property details update failed"})
		return
	}
	UpdatedCommercialPropertyForSale := commercialRepo.GetCommercialPropertyForSaleWithAllAssociationsById(c.Param("id"))
	if UpdatedCommercialPropertyForSale == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.CommercialPropertyForSaleResponse(*UpdatedCommercialPropertyForSale)})
}

func GetAllCommercialForSaleProperties(c *gin.Context) {
	commercialPropertiesForSale := commercialRepo.GetAllCommercialPropertiesForSale()
	responseList := []commercialDtos.CommercialForSalePropertyWithManagerResponseDto{}
	if len(commercialPropertiesForSale) > 0 {
		for i := 0; i < len(commercialPropertiesForSale); i++ {
			responseItem := propertyUtilities.CommercialPropertyForSaleWithManagerResponse(commercialPropertiesForSale[i])
			responseList = append(responseList, responseItem)
		}
		c.JSON(http.StatusOK, gin.H{
			"response": responseList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"response": responseList,
	})
}

func GetCommercialPropertyForSaleById(c *gin.Context) {
	commercialPropertyForSale := commercialRepo.GetCommercialPropertyForSaleWithAllAssociationsById(c.Param("id"))
	if commercialPropertyForSale == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.CommercialPropertyForSaleWithManagerResponse(*commercialPropertyForSale)})
}

func GetManagerCommercialPropertiesForSaleByManagerId(c *gin.Context) {
	properties := commercialRepo.GetManagerCommercialPropertiesForSaleByManagerId(c.Param("id"))
	propertiesResponse := []commercialDtos.CommercialForSalePropertyResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			propertyResponse := propertyUtilities.CommercialPropertyForSaleResponse(properties[i])
			propertiesResponse = append(propertiesResponse, propertyResponse)
		}
	}
	c.JSON(http.StatusOK, gin.H{"response": propertiesResponse})
}

func DeleteCommercialPropertyForSaleById(c *gin.Context) {
	isPropertyDeleted := commercialRepo.DeleteCommercialPropertyForSaleById(c.Param("id"))
	if !isPropertyDeleted {
		c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"response": "your property was successfully deleted"})
		return
	}
}
