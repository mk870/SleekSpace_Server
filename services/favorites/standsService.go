package favorites

import (
	"net/http"

	favoritesDtos "SleekSpace/dtos/favorites"
	standDtos "SleekSpace/dtos/property/stand"
	favoritesRepo "SleekSpace/repositories/favorites"
	userRepo "SleekSpace/repositories/user"
	constantsUtilities "SleekSpace/utilities/constants"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetFavoriteStandProperties(c *gin.Context) {
	user := userRepo.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	properties := favoritesRepo.GetFavoriteStandProperties(user.FavoriteStands)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "your don't have stand property favorites"})
		return
	}
	response := []standDtos.StandWithManagerResponseDTO{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			response = append(response, propertyUtilities.PropertyStandWithManagerResponse(properties[i]))
		}
	}

	presentFavouriteStandProperties := []int{}
	for i := 0; i < len(response); i++ {
		presentFavouriteStandProperties = append(presentFavouriteStandProperties, response[i].Id)
	}

	deletedFavouriteProperties := propertyUtilities.ReturnDeletedFavoriteProperties(presentFavouriteStandProperties, user.FavoriteStands)

	c.JSON(http.StatusOK, gin.H{
		"deleted":   deletedFavouriteProperties,
		"available": response,
	})
}

func UpdateStandPropertyFavourites(c *gin.Context) {
	var standPropertiesIds favoritesDtos.FavouritesPropertyIds
	validateModelFields := validator.New()
	c.BindJSON(&standPropertiesIds)

	modelFieldsValidationError := validateModelFields.Struct(standPropertiesIds)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	user := userRepo.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	user.FavoriteStands = standPropertiesIds.Ids
	isUpdated := userRepo.SaveUserUpdate(user)
	if !isUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your stand property favorites"})
		return
	}
	properties := favoritesRepo.GetFavoriteStandProperties(standPropertiesIds.Ids)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "your don't have stand property favorites"})
		return
	}
	response := []standDtos.StandWithManagerResponseDTO{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			response = append(response, propertyUtilities.PropertyStandWithManagerResponse(properties[i]))
		}
	}

	presentFavouriteStandProperties := []int{}
	for i := 0; i < len(response); i++ {
		presentFavouriteStandProperties = append(presentFavouriteStandProperties, response[i].Id)
	}

	deletedFavouriteProperties := propertyUtilities.ReturnDeletedFavoriteProperties(presentFavouriteStandProperties, standPropertiesIds.Ids)

	c.JSON(http.StatusOK, gin.H{
		"deleted":   deletedFavouriteProperties,
		"available": response,
	})
}
