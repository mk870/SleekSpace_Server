package property

import (
	"net/http"

	commercialRepo "SleekSpace/repositories/property/commercial"
	landRepo "SleekSpace/repositories/property/land"
	residentialRepo "SleekSpace/repositories/property/residential"
	standRepo "SleekSpace/repositories/property/stand"
	generalUtilities "SleekSpace/utilities/funcs/general"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
)

func GetPropertyTypeById(c *gin.Context, propertyType string, propertyId int) {
	if propertyType == "stand" {
		stand := standRepo.GetStandWithAllAssociationsById(generalUtilities.ConvertIntToString(propertyId))
		if stand == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "this stand does not exist"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.PropertyStandResponse(*stand)})
		return
	} else if propertyType == "land" {
		land := landRepo.GetLandPropertyForSaleWithAllAssociationsById(generalUtilities.ConvertIntToString(propertyId))
		if land == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "this land does not exist"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.LandPropertyResponse(*land)})
		return
	} else if propertyType == "residentialRentalProperty" {
		property := residentialRepo.GetResidentialRentalPropertyWithAllAssociationsById(generalUtilities.ConvertIntToString(propertyId))
		if property == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.ResidentialRentalPropertyResponse(*property)})
		return
	} else if propertyType == "residentialPropertyForSale" {
		property := residentialRepo.GetResidentialPropertyForSaleWithAllAssociationsById(generalUtilities.ConvertIntToString(propertyId))
		if property == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.ResidentialForSalePropertyResponse(*property)})
		return
	} else if propertyType == "commercialRentalProperty" {
		property := commercialRepo.GetCommercialRentalPropertyWithAllAssociationsById(generalUtilities.ConvertIntToString(propertyId))
		if property == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.CommercialPropertyForRentResponse(*property)})
		return
	} else {
		property := commercialRepo.GetCommercialPropertyForSaleWithAllAssociationsById(generalUtilities.ConvertIntToString(propertyId))
		if property == nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "this property does not exist"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.CommercialPropertyForSaleResponse(*property)})
		return
	}
}
