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

func GetFavoriteResidentialForSaleProperties(c *gin.Context) {
	user := userRepo.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	properties := favoritesRepo.GetFavoriteResidentialPropertiesForSale(user.FavoriteResidentialForSaleProperties)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "your don't have favorite residential for sale properties"})
		return
	}
	response := []residentialDtos.ResidentialPropertyForSaleResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			response = append(response, propertyUtilities.ResidentialForSalePropertyResponse(properties[i]))
		}
	}

	presentFavouriteResidentialForSaleProperties := []int{}
	for i := 0; i < len(response); i++ {
		presentFavouriteResidentialForSaleProperties = append(presentFavouriteResidentialForSaleProperties, response[i].Id)
	}

	deletedFavouriteProperties := propertyUtilities.ReturnDeletedFavoriteProperties(presentFavouriteResidentialForSaleProperties, user.FavoriteResidentialForSaleProperties)

	c.JSON(http.StatusOK, gin.H{
		"deleted":   deletedFavouriteProperties,
		"available": response,
	})
}

func UpdateFavouritesResidentialForSaleProperties(c *gin.Context) {
	var residentialForSalePropertiesIds favoritesDtos.FavouritesPropertyIds
	validateModelFields := validator.New()
	c.BindJSON(&residentialForSalePropertiesIds)

	modelFieldsValidationError := validateModelFields.Struct(residentialForSalePropertiesIds)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	user := userRepo.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	user.FavoriteResidentialForSaleProperties = residentialForSalePropertiesIds.Ids
	isUpdated := userRepo.SaveUserUpdate(user)
	if !isUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your property favorites"})
		return
	}
	properties := favoritesRepo.GetFavoriteResidentialPropertiesForSale(residentialForSalePropertiesIds.Ids)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "your don't have favorite residential for sale properties"})
		return
	}
	response := []residentialDtos.ResidentialPropertyForSaleResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			response = append(response, propertyUtilities.ResidentialForSalePropertyResponse(properties[i]))
		}
	}

	presentFavouriteResidentialForSaleProperties := []int{}
	for i := 0; i < len(response); i++ {
		presentFavouriteResidentialForSaleProperties = append(presentFavouriteResidentialForSaleProperties, response[i].Id)
	}

	deletedFavouriteProperties := propertyUtilities.ReturnDeletedFavoriteProperties(presentFavouriteResidentialForSaleProperties, residentialForSalePropertiesIds.Ids)

	c.JSON(http.StatusOK, gin.H{
		"deleted":   deletedFavouriteProperties,
		"available": response,
	})
}
