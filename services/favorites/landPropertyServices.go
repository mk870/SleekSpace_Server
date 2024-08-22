package favorites

import (
	"net/http"

	favoritesDtos "SleekSpace/dtos/favorites"
	landDtos "SleekSpace/dtos/property/land"
	favoritesRepo "SleekSpace/repositories/favorites"
	userRepo "SleekSpace/repositories/user"
	constantsUtilities "SleekSpace/utilities/constants"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetFavoriteLandProperties(c *gin.Context) {
	user := userRepo.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	properties := favoritesRepo.GetFavoriteLandProperties(user.FavoriteLandProperties)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "your don't have land property favorites"})
		return
	}
	response := []landDtos.LandForSalePropertyResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			response = append(response, propertyUtilities.LandPropertyResponse(properties[i]))
		}
	}

	presentFavouriteLandProperties := []int{}
	for i := 0; i < len(response); i++ {
		presentFavouriteLandProperties = append(presentFavouriteLandProperties, response[i].Id)
	}

	deletedFavouriteProperties := propertyUtilities.ReturnDeletedFavoriteProperties(presentFavouriteLandProperties, user.FavoriteLandProperties)

	c.JSON(http.StatusOK, gin.H{
		"deleted":   deletedFavouriteProperties,
		"available": response,
	})
}

func UpdateLandPropertyFavourites(c *gin.Context) {
	var landPropertiesIds favoritesDtos.FavouritesPropertyIds
	validateModelFields := validator.New()
	c.BindJSON(&landPropertiesIds)

	modelFieldsValidationError := validateModelFields.Struct(landPropertiesIds)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	user := userRepo.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	user.FavoriteLandProperties = landPropertiesIds.Ids
	isUpdated := userRepo.SaveUserUpdate(user)
	if !isUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your land property favorites"})
		return
	}
	properties := favoritesRepo.GetFavoriteLandProperties(landPropertiesIds.Ids)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "your don't have land property favorites"})
		return
	}
	response := []landDtos.LandForSalePropertyResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			response = append(response, propertyUtilities.LandPropertyResponse(properties[i]))
		}
	}

	presentFavouriteLandProperties := []int{}
	for i := 0; i < len(response); i++ {
		presentFavouriteLandProperties = append(presentFavouriteLandProperties, response[i].Id)
	}

	deletedFavouriteProperties := propertyUtilities.ReturnDeletedFavoriteProperties(presentFavouriteLandProperties, landPropertiesIds.Ids)

	c.JSON(http.StatusOK, gin.H{
		"deleted":   deletedFavouriteProperties,
		"available": response,
	})
}
