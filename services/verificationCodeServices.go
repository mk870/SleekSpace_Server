package services

import (
	"net/http"
	"strconv"
	"time"

	"SleekSpace/dtos"
	"SleekSpace/repositories"
	"SleekSpace/tokens"
	"SleekSpace/utilities"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func VerifyCodeForRegistration(c *gin.Context) {
	var verificationInfo = dtos.VerificationDTO{}
	validateModelFields := validator.New()
	c.BindJSON(&verificationInfo)
	modelFieldsValidationError := validateModelFields.Struct(verificationInfo)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}
	storedVerificationCode := repositories.GetVerificationCodeById(utilities.ConvertIntToString(verificationInfo.UserId))

	if storedVerificationCode.ExpiryDate.Unix() < time.Now().Local().Unix() {
		isUserDeleted := repositories.DeleteUserById(strconv.Itoa(storedVerificationCode.UserId))
		if !isUserDeleted {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "this user does not exist",
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "this verification code has expired please signup again",
			})
			return
		}
	}
	if verificationInfo.VerificationCode != storedVerificationCode.Code {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong verification code, please try again",
		})
		return
	}
	user := repositories.GetUserById(strconv.Itoa(storedVerificationCode.UserId))
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	user.IsActive = true
	isUpdated := repositories.SaveUserUpdate(user)
	if isUpdated {
		accessToken := tokens.GenerateAccessToken(user.GivenName, user.Email, user.Id)
		if accessToken == "failed" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "could not generate your access token",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"accessToken": accessToken,
			"id":          user.Id,
			"email":       user.Email,
			"givenName":   user.GivenName,
			"familyName":  user.FamilyName,
		})
		return
	}
}

func CreateVerificationCode(c *gin.Context) {
	userEmail := dtos.CreateVerificationCodeDTO{}
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

func VerifyCodeForSecurity(c *gin.Context) {
	var verificationInfo = dtos.VerificationDTO{}
	validateModelFields := validator.New()
	c.BindJSON(&verificationInfo)
	modelFieldsValidationError := validateModelFields.Struct(verificationInfo)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}
	storedVerificationCode := repositories.GetVerificationCodeById(utilities.ConvertIntToString(verificationInfo.UserId))

	if storedVerificationCode.ExpiryDate.Unix() < time.Now().Local().Unix() {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "this verification code has expired please signup again",
		})
		return
	}
	if verificationInfo.VerificationCode != storedVerificationCode.Code {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong verification code, please try again",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"userId": storedVerificationCode.UserId,
	})
}

func ResendVerificationCode(c *gin.Context) {
	userId := c.Param("id")
	user := repositories.GetUserById(userId)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this user does not exist"})
		return
	}
	verificationCode := repositories.GetVerificationCodeById(userId)
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
		"response": "please check your email for verification code",
	})
}
