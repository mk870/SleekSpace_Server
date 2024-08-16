package services

import (
	"net/http"

	userDtos "SleekSpace/dtos/user"
	userModels "SleekSpace/models/user"
	userRepo "SleekSpace/repositories/user"
	"SleekSpace/tokens"
	"SleekSpace/utilities"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func Logout(c *gin.Context) {
	user := c.MustGet("user").(*userModels.User)
	loggedInUser := userRepo.GetUserByEmail(user.Email)
	if loggedInUser == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get user information",
		})
		return
	}
	isRefreshTokenDeleted := userRepo.SaveUserUpdate(loggedInUser)
	if isRefreshTokenDeleted {
		c.JSON(http.StatusOK, "you have logged out successfully")
		return
	}
}

func Login(c *gin.Context) {
	var loginData = userDtos.LoginDTO{}
	validateModelFields := validator.New()
	c.BindJSON(&loginData)
	modelFieldsValidationError := validateModelFields.Struct(loginData)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	user := userRepo.GetUserByEmail(loginData.Email)
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "this account does not exist"})
		return
	}
	if user.IsSocialsAuthenticated {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please login using socials since you signed up with it"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong password or email"})
		return
	}

	accessToken := tokens.GenerateAccessToken(user.GivenName, user.Email, user.Id)
	if accessToken == "failed" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create access token",
		})
		return
	}
	user.AccessToken = accessToken
	isUpdated := userRepo.SaveUserUpdate(user)
	if !isUpdated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update the accessToken",
		})
		return
	}
	response := utilities.UserResponseMapper(user, accessToken)
	c.JSON(http.StatusOK, gin.H{"response": response})
}
