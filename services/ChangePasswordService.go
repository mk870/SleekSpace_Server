package services

import (
	"net/http"
	"time"

	"SleekSpace/dtos"
	"SleekSpace/repositories"
	"SleekSpace/utilities"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func ChangePassword(c *gin.Context) {
	userEmail := dtos.ChangePasswordDTO{}
	validateModelFields := validator.New()
	c.BindJSON(&userEmail)
	modelFieldsValidationError := validateModelFields.Struct(userEmail)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}
	user := repositories.GetUserByEmail(userEmail.Email)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	verificationCode := repositories.GetVerificationCodeById(utilities.ConvertIntToString(user.Id))
	verificationCode.Code = utilities.GenerateVerificationCode()
	verificationCode.ExpiryDate = time.Now().Add(time.Minute * 15)
	isVerificationCodeUpdated := repositories.UpdateVerificationCode(&verificationCode)
	if !isVerificationCodeUpdated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not generate another code",
		})
		return
	}
	isEmailSent := SendVerificationCodeEmail(user.Email, user.GivenName, utilities.ConvertIntToString(verificationCode.Code))
	if !isEmailSent {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send verification email"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"userId": user.Id,
	})
}

func UpdatePasswordAndCodeVerification(c *gin.Context) {
	var newPasswordData = dtos.NewPasswordDTO{}
	validateModelFields := validator.New()
	c.BindJSON(&newPasswordData)
	modelFieldsValidationError := validateModelFields.Struct(newPasswordData)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}
	storedVerificationCode := repositories.GetVerificationCodeById(utilities.ConvertIntToString(newPasswordData.UserId))

	if storedVerificationCode.ExpiryDate.Unix() < time.Now().Local().Unix() {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "this verification code has expired try again",
		})
		return
	}

	if newPasswordData.VerificationCode != storedVerificationCode.Code {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong verification code, try again",
		})
		return
	}

	user := repositories.GetUserById(utilities.ConvertIntToString(newPasswordData.UserId))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	if user.IsSocialsAuthenticated {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please change socials password since you signed up with it"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPasswordData.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)
	updateResult := repositories.SaveUserUpdate(user)
	if !updateResult {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "password update failed please try again later"})
	}
	c.JSON(http.StatusOK, gin.H{"response": "password update successful"})
}
