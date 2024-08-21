package admin

import (
	"fmt"
	"net/http"
	"time"

	commercialDtos "SleekSpace/dtos/property/commercial"
	landDtos "SleekSpace/dtos/property/land"
	propertyLocationDtos "SleekSpace/dtos/property/location"
	residentialDtos "SleekSpace/dtos/property/residential"
	standDtos "SleekSpace/dtos/property/stand"
	managerRepo "SleekSpace/repositories/manager"
	commercialRepo "SleekSpace/repositories/property/commercial"
	propertyInsightsRepo "SleekSpace/repositories/property/insights"
	landRepo "SleekSpace/repositories/property/land"
	propertyLocationRepo "SleekSpace/repositories/property/location"
	propertyMediaRepo "SleekSpace/repositories/property/media"
	residentialRepo "SleekSpace/repositories/property/residential"
	standRepo "SleekSpace/repositories/property/stand"
	userRepo "SleekSpace/repositories/user"
	generalUtilities "SleekSpace/utilities/funcs/general"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
)

func GetLocationById(c *gin.Context) {
	id := c.Param("id")
	location := userRepo.GetLocationById(generalUtilities.ConvertStringToInt(id))
	c.JSON(http.StatusOK, gin.H{"response": location})
}

func apiSim() <-chan string {
	result := make(chan string)
	go func() {
		defer close(result)
		time.Sleep(time.Second * 5)
		result <- "im done"
	}()
	return result
}

func GetInfo(c *gin.Context) {
	info := <-apiSim()
	fmt.Println("ran aftrer")
	c.JSON(http.StatusOK, gin.H{"response": info})

}

func GetAllUsersLocations(c *gin.Context) {
	codes := userRepo.GetAllUsersLocations()
	c.JSON(http.StatusOK, gin.H{
		"response": codes,
	})
}

func GetAllUsers(c *gin.Context) {
	users := userRepo.GetUsers()
	c.JSON(http.StatusOK, gin.H{
		"response": users,
	})
}

func GetAllManagersProfilePictures(c *gin.Context) {
	pictures := managerRepo.GetAllManagersProfilePictures()
	c.JSON(http.StatusOK, gin.H{
		"response": pictures,
	})
}

func GetVerificationCodeById(c *gin.Context) {
	id := c.Param("id")
	code := userRepo.GetVerificationCodeById(id)
	c.JSON(http.StatusOK, gin.H{
		"response": code,
	})
}

func GetAllVerificationCodes(c *gin.Context) {
	codes := userRepo.AllVerificationCodes()
	c.JSON(http.StatusOK, gin.H{
		"response": codes,
	})
}

func DeleteVerificationCode(c *gin.Context) {
	id := c.Param("id")
	isDeleted := userRepo.DeleteVerficationCode(generalUtilities.ConvertStringToInt(id))
	if isDeleted {
		c.JSON(http.StatusOK, gin.H{
			"response": "code deleted",
		})
	}
}

func GetUserContacts(c *gin.Context) {
	numbers := userRepo.GetAllUsersContactNumbers()
	c.JSON(http.StatusOK, gin.H{
		"response": numbers,
	})
}

func GetAllManagersContacts(c *gin.Context) {
	numbers := managerRepo.GetAllManagersContacts()
	c.JSON(http.StatusOK, gin.H{
		"response": numbers,
	})
}

func GetAllManagers(c *gin.Context) {
	managers := managerRepo.GetAllManagers()
	c.JSON(http.StatusOK, gin.H{
		"response": managers,
	})
}

func GetAllStands(c *gin.Context) {
	stands := standRepo.GetAllStands()
	responseList := []standDtos.StandResponseDTO{}
	if len(stands) > 0 {
		for i := 0; i < len(stands); i++ {
			responseItem := propertyUtilities.PropertyStandResponse(stands[i])
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

func GetAllLandProperties(c *gin.Context) {
	landProperties := landRepo.GetAllLandPropertiesForSale()
	responseList := []landDtos.LandForSalePropertyResponseDto{}
	if len(landProperties) > 0 {
		for i := 0; i < len(landProperties); i++ {
			responseItem := propertyUtilities.LandPropertyResponse(landProperties[i])
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

func GetAllResidentialForSaleProperties(c *gin.Context) {
	residentialPropertiesForSale := residentialRepo.GetAllResidentialPropertiesForSale()
	responseList := []residentialDtos.ResidentialPropertyForSaleResponseDto{}
	if len(residentialPropertiesForSale) > 0 {
		for i := 0; i < len(residentialPropertiesForSale); i++ {
			responseItem := propertyUtilities.ResidentialForSalePropertyResponse(residentialPropertiesForSale[i])
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

func GetAllResidentialRentalProperties(c *gin.Context) {
	residentialRentalProperties := residentialRepo.GetAllResidentialRentalProperties()
	responseList := []residentialDtos.ResidentialPropertyForRentResponseDto{}
	if len(residentialRentalProperties) > 0 {
		for i := 0; i < len(residentialRentalProperties); i++ {
			responseItem := propertyUtilities.ResidentialRentalPropertyResponse(residentialRentalProperties[i])
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

func GetAllCommercialRentalProperties(c *gin.Context) {
	commercialRentalProperties := commercialRepo.GetAllCommercialRentalProperties()
	responseList := []commercialDtos.CommercialForRentPropertyResponseDto{}
	if len(commercialRentalProperties) > 0 {
		for i := 0; i < len(commercialRentalProperties); i++ {
			responseItem := propertyUtilities.CommercialPropertyForRentResponse(commercialRentalProperties[i])
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

func GetAllCommercialForSaleProperties(c *gin.Context) {
	commercialPropertiesForSale := commercialRepo.GetAllCommercialPropertiesForSale()
	responseList := []commercialDtos.CommercialForSalePropertyResponseDto{}
	if len(commercialPropertiesForSale) > 0 {
		for i := 0; i < len(commercialPropertiesForSale); i++ {
			responseItem := propertyUtilities.CommercialPropertyForSaleResponse(commercialPropertiesForSale[i])
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

func GetAllPropertiesImagesOrVideos(c *gin.Context) {
	imagesOrVideosList := propertyMediaRepo.GetAllPropertiesImagesOrVideos()
	c.JSON(http.StatusOK, gin.H{
		"response": imagesOrVideosList,
	})
}

func GetAllPropertiesInsights(c *gin.Context) {
	insightsList := propertyInsightsRepo.GetAllPropertiesInsights()
	c.JSON(http.StatusOK, gin.H{
		"response": insightsList,
	})
}

func GetAllPropertiesLocation(c *gin.Context) {
	propertyLocations := propertyLocationRepo.GetAllPropertyLocations()
	responseList := []propertyLocationDtos.PropertyLocationUpdateAndResponseDto{}
	if len(propertyLocations) > 0 {
		for i := 0; i < len(propertyLocations); i++ {
			responseItem := propertyUtilities.PropertyLocationResponse(propertyLocations[i])
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
