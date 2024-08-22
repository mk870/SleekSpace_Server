package favorites

import (
	"SleekSpace/middleware"
	favoritesService "SleekSpace/services/favorites"

	"github.com/gin-gonic/gin"
)

func PropertyFavoritesRoutes(router *gin.Engine) {
	routes := router.Group("/favorites/property")
	{
		routes.GET("/land/:id", middleware.AuthValidator, favoritesService.GetFavoriteLandProperties)
		routes.PUT("/land/:id", middleware.AuthValidator, favoritesService.UpdateLandPropertyFavourites)
		routes.GET("/stand/:id", middleware.AuthValidator, favoritesService.GetFavoriteStandProperties)
		routes.PUT("/stand/:id", middleware.AuthValidator, favoritesService.UpdateStandPropertyFavourites)
		routes.GET("/commercial/onsale/:id", middleware.AuthValidator, favoritesService.GetFavoriteCommercialForSaleProperties)
		routes.PUT("/commercial/onsale/:id", middleware.AuthValidator, favoritesService.UpdateFavouritesCommercialForSaleProperties)
		routes.GET("/commercial/rentals/:id", middleware.AuthValidator, favoritesService.GetFavoriteCommercialRentalProperties)
		routes.PUT("/commercial/rentals/:id", middleware.AuthValidator, favoritesService.UpdateFavouritesCommercialRentalProperties)
		routes.GET("/residential/onsale/:id", middleware.AuthValidator, favoritesService.GetFavoriteResidentialForSaleProperties)
		routes.PUT("/residential/onsale/:id", middleware.AuthValidator, favoritesService.UpdateFavouritesResidentialForSaleProperties)
		routes.GET("/residential/rentals/:id", middleware.AuthValidator, favoritesService.GetFavoriteResidentialRentalProperties)
		routes.PUT("/residential/rentals/:id", middleware.AuthValidator, favoritesService.UpdateFavouritesResidentialRentalProperties)
	}
}
