package favorites

import (
	"net/http"

	favoritesDtos "SleekSpace/dtos/favorites"
	residentialDtos "SleekSpace/dtos/property/residential"
	favoritesRepo "SleekSpace/repositories/favorites"
	userRepo "SleekSpace/repositories/user"
	constantsUtilities "SleekSpace/utilities/constants"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetFavoriteResidentialRentalProperties(c *gin.Context) {
	user := userRepo.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	properties := favoritesRepo.GetFavoriteResidentialRentalProperties(user.FavoriteResidentialRentalProperties)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "your don't have favorite residential rental properties"})
		return
	}
	response := []residentialDtos.ResidentialPropertyForRentResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			response = append(response, propertyUtilities.ResidentialRentalPropertyResponse(properties[i]))
		}
	}

	presentFavouriteResidentialRentalProperties := []int{}
	for i := 0; i < len(response); i++ {
		presentFavouriteResidentialRentalProperties = append(presentFavouriteResidentialRentalProperties, response[i].Id)
	}

	deletedFavouriteProperties := propertyUtilities.ReturnDeletedFavoriteProperties(presentFavouriteResidentialRentalProperties, user.FavoriteResidentialRentalProperties)

	c.JSON(http.StatusOK, gin.H{
		"deleted":   deletedFavouriteProperties,
		"available": response,
	})
}

func UpdateFavouritesResidentialRentalProperties(c *gin.Context) {
	var residentialRentalPropertiesIds favoritesDtos.FavouritesPropertyIds
	validateModelFields := validator.New()
	c.BindJSON(&residentialRentalPropertiesIds)

	modelFieldsValidationError := validateModelFields.Struct(residentialRentalPropertiesIds)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	user := userRepo.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	user.FavoriteResidentialRentalProperties = residentialRentalPropertiesIds.Ids
	isUpdated := userRepo.SaveUserUpdate(user)
	if !isUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your property favorites"})
		return
	}
	properties := favoritesRepo.GetFavoriteResidentialRentalProperties(residentialRentalPropertiesIds.Ids)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "your don't have favorite residential rental properties"})
		return
	}
	response := []residentialDtos.ResidentialPropertyForRentResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			response = append(response, propertyUtilities.ResidentialRentalPropertyResponse(properties[i]))
		}
	}

	presentFavouriteResidentialRentalProperties := []int{}
	for i := 0; i < len(response); i++ {
		presentFavouriteResidentialRentalProperties = append(presentFavouriteResidentialRentalProperties, response[i].Id)
	}

	deletedFavouriteProperties := propertyUtilities.ReturnDeletedFavoriteProperties(presentFavouriteResidentialRentalProperties, residentialRentalPropertiesIds.Ids)

	c.JSON(http.StatusOK, gin.H{
		"deleted":   deletedFavouriteProperties,
		"available": response,
	})
}
