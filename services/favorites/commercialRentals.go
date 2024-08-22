package favorites

import (
	"net/http"

	favoritesDtos "SleekSpace/dtos/favorites"
	commercialDtos "SleekSpace/dtos/property/commercial"
	favoritesRepo "SleekSpace/repositories/favorites"
	userRepo "SleekSpace/repositories/user"
	constantsUtilities "SleekSpace/utilities/constants"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetFavoriteCommercialRentalProperties(c *gin.Context) {
	user := userRepo.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	properties := favoritesRepo.GetFavoriteCommercialRentalProperties(user.FavoriteCommercialRentalProperties)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "your don't have favorite commercial rental properties"})
		return
	}
	response := []commercialDtos.CommercialForRentPropertyResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			response = append(response, propertyUtilities.CommercialPropertyForRentResponse(properties[i]))
		}
	}

	presentFavouriteCommercialRentalProperties := []int{}
	for i := 0; i < len(response); i++ {
		presentFavouriteCommercialRentalProperties = append(presentFavouriteCommercialRentalProperties, response[i].Id)
	}

	deletedFavouriteProperties := propertyUtilities.ReturnDeletedFavoriteProperties(presentFavouriteCommercialRentalProperties, user.FavoriteCommercialRentalProperties)

	c.JSON(http.StatusOK, gin.H{
		"deleted":   deletedFavouriteProperties,
		"available": response,
	})
}

func UpdateFavouritesCommercialRentalProperties(c *gin.Context) {
	var commercialRentalPropertiesIds favoritesDtos.FavouritesPropertyIds
	validateModelFields := validator.New()
	c.BindJSON(&commercialRentalPropertiesIds)

	modelFieldsValidationError := validateModelFields.Struct(commercialRentalPropertiesIds)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	user := userRepo.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	user.FavoriteCommercialRentalProperties = commercialRentalPropertiesIds.Ids
	isUpdated := userRepo.SaveUserUpdate(user)
	if !isUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your property favorites"})
		return
	}
	properties := favoritesRepo.GetFavoriteCommercialRentalProperties(commercialRentalPropertiesIds.Ids)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "your don't have favorite commercial rental properties"})
		return
	}
	response := []commercialDtos.CommercialForRentPropertyResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			response = append(response, propertyUtilities.CommercialPropertyForRentResponse(properties[i]))
		}
	}

	presentFavouriteCommercialRentalProperties := []int{}
	for i := 0; i < len(response); i++ {
		presentFavouriteCommercialRentalProperties = append(presentFavouriteCommercialRentalProperties, response[i].Id)
	}

	deletedFavouriteProperties := propertyUtilities.ReturnDeletedFavoriteProperties(presentFavouriteCommercialRentalProperties, commercialRentalPropertiesIds.Ids)

	c.JSON(http.StatusOK, gin.H{
		"deleted":   deletedFavouriteProperties,
		"available": response,
	})
}
