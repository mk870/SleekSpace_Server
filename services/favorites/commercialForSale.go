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

func GetFavoriteCommercialForSaleProperties(c *gin.Context) {
	user := userRepo.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	properties := favoritesRepo.GetFavoriteCommercialPropertiesForSale(user.FavoriteCommercialForSaleProperties)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "your don't have favorite commercial for sale properties"})
		return
	}
	response := []commercialDtos.CommercialForSalePropertyResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			response = append(response, propertyUtilities.CommercialPropertyForSaleResponse(properties[i]))
		}
	}

	presentFavouriteCommercialForSaleProperties := []int{}
	for i := 0; i < len(response); i++ {
		presentFavouriteCommercialForSaleProperties = append(presentFavouriteCommercialForSaleProperties, response[i].Id)
	}

	deletedFavouriteProperties := propertyUtilities.ReturnDeletedFavoriteProperties(presentFavouriteCommercialForSaleProperties, user.FavoriteCommercialForSaleProperties)

	c.JSON(http.StatusOK, gin.H{
		"deleted":   deletedFavouriteProperties,
		"available": response,
	})
}

func UpdateFavouritesCommercialForSaleProperties(c *gin.Context) {
	var commercialForSalePropertiesIds favoritesDtos.FavouritesPropertyIds
	validateModelFields := validator.New()
	c.BindJSON(&commercialForSalePropertiesIds)

	modelFieldsValidationError := validateModelFields.Struct(commercialForSalePropertiesIds)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	user := userRepo.GetUserById(c.Param("id"))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": constantsUtilities.NoUserError})
		return
	}
	user.FavoriteCommercialForSaleProperties = commercialForSalePropertiesIds.Ids
	isUpdated := userRepo.SaveUserUpdate(user)
	if !isUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your property favorites"})
		return
	}
	properties := favoritesRepo.GetFavoriteCommercialPropertiesForSale(commercialForSalePropertiesIds.Ids)
	if properties == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "your don't have favorite commercial for sale properties"})
		return
	}
	response := []commercialDtos.CommercialForSalePropertyResponseDto{}
	if len(properties) > 0 {
		for i := 0; i < len(properties); i++ {
			response = append(response, propertyUtilities.CommercialPropertyForSaleResponse(properties[i]))
		}
	}

	presentFavouriteCommercialForSaleProperties := []int{}
	for i := 0; i < len(response); i++ {
		presentFavouriteCommercialForSaleProperties = append(presentFavouriteCommercialForSaleProperties, response[i].Id)
	}

	deletedFavouriteProperties := propertyUtilities.ReturnDeletedFavoriteProperties(presentFavouriteCommercialForSaleProperties, commercialForSalePropertiesIds.Ids)

	c.JSON(http.StatusOK, gin.H{
		"deleted":   deletedFavouriteProperties,
		"available": response,
	})
}
