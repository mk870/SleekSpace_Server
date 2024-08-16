package main

import (
	adminController "SleekSpace/controllers/admin"
	authController "SleekSpace/controllers/auth"
	externalApiCallsController "SleekSpace/controllers/externalApiCalls"
	managerController "SleekSpace/controllers/manager"
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

	adminController.GetVerificationCodeById(router)
	adminController.DeleteVerificationCode(router)
	adminController.GetAllManagers(router)
	adminController.GetAllManagersContacts(router)
	adminController.GetAllUsersLocations(router)
	adminController.GetLocationById(router)
	adminController.GetUserContacts(router)
	adminController.GetAllVerificationCodes(router)
	adminController.GetAllManagersProfilePictures(router)

	router.Run()
}
