package main

import (
	adminController "SleekSpace/controllers/admin"
	authController "SleekSpace/controllers/auth"
	externalApiCallsController "SleekSpace/controllers/externalApiCalls"
	managerController "SleekSpace/controllers/manager"
	insightsController "SleekSpace/controllers/property/insights"
	locationController "SleekSpace/controllers/property/location"
	mediaController "SleekSpace/controllers/property/media"
	standsController "SleekSpace/controllers/property/stand"
	userController "SleekSpace/controllers/user"
	"SleekSpace/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization", "token", "User-Agent", "Accept")
	router.Use(cors.New(config))
	db.Connect()

	userController.UpdateContactNumbers(router)
	userController.CreateContact(router)
	userController.CreateLocation(router)
	userController.UpdateLocation(router)
	userController.GetUsers(router)
	userController.UpdateUser(router)
	userController.GetUser(router)
	userController.GetUserByEmail(router)
	userController.DeleteUser(router)
	userController.CreateUserProfilePicture(router)
	userController.UpdateUserProfilePicture(router)

	authController.Registration(router)
	authController.Login(router)
	authController.LoginOut(router)
	authController.RegistrationVerificationCodeValidation(router)
	authController.SecurityVerificationCodeValidation(router)
	authController.ResendVerificationCode(router)
	authController.CreateVerificationCode(router)
	authController.PasswordUpdate(router)

	externalApiCallsController.LocationAutoComplete(router)
	externalApiCallsController.LocationReverseGeoCoding(router)

	managerController.CreateManager(router)
	managerController.GetManagerById(router)
	managerController.GetManagerByUserId(router)
	managerController.DeleteManager(router)
	managerController.UpdateManager(router)
	managerController.UpdateManagerProfilePicture(router)
	managerController.UpdateManagerContactNumbers(router)
	managerController.GetManagerStandsByManagerId(router)

	adminController.GetVerificationCodeById(router)
	adminController.DeleteVerificationCode(router)
	adminController.GetAllManagers(router)
	adminController.GetAllManagersContacts(router)
	adminController.GetAllUsersLocations(router)
	adminController.GetLocationById(router)
	adminController.GetUserContacts(router)
	adminController.GetAllVerificationCodes(router)
	adminController.GetAllManagersProfilePictures(router)
	adminController.ApiSim(router)
	adminController.GetAllStands(router)
	adminController.GetAllPropertiesImagesOrVideos(router)
	adminController.GetAllPropertiesLocations(router)
	adminController.GetAllPropertiesInsights(router)

	standsController.CreateStandForSale(router)
	standsController.DeleteStandById(router)
	standsController.UpdateStandDetails(router)
	standsController.GetStandById(router)

	mediaController.CreatePropertyImageOrVideo(router)
	mediaController.DeletePropertyImageOrVideo(router)
	mediaController.UpdatePropertyImageOrVideo(router)
	mediaController.GetPropertyImageOrVideoById(router)

	locationController.GetPropertyLocationById(router)
	locationController.UpdatePropertyLocation(router)

	insightsController.GetPropertyInsightsById(router)
	insightsController.GetPropertyInsightsByPropertyId(router)
	insightsController.UpdatePropertyInsights(router)

	router.Run()
}
