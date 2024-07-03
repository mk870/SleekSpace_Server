package services

import (
	"net/http"

	"SleekSpace/dtos"
	"SleekSpace/models"
	"SleekSpace/repositories"
	"SleekSpace/utilities"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func Registration(c *gin.Context) {
	var userRegistartionDTO dtos.NativeUserRegistrationDTO
	validateModelFields := validator.New()
	c.BindJSON(&userRegistartionDTO)

	modelFieldsValidationError := validateModelFields.Struct(userRegistartionDTO)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRegistartionDTO.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	newUser := models.User{
		GivenName:              userRegistartionDTO.GivenName,
		FamilyName:             userRegistartionDTO.FamilyName,
		Password:               string(hashedPassword),
		Email:                  userRegistartionDTO.Email,
		IsSocialsAuthenticated: false,
		RegistrationCode: models.VerificationCode{
			Code:       utilities.GenerateVerificationCode(),
			ExpiryDate: utilities.GenerateVerificationGracePeriod(),
		},
	}

	isUserCreated := repositories.CreateUser(&newUser)
	if !isUserCreated {
		c.JSON(http.StatusForbidden, gin.H{"error": "this email already exists"})
		return
	}

	isVerificationEmailSent := SendVerificationCodeEmail(userRegistartionDTO.Email, userRegistartionDTO.GivenName, utilities.ConvertIntToString(newUser.RegistrationCode.Code))
	if !isVerificationEmailSent {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send verification email"})
		return
	}
	createdUser := repositories.GetUserByEmail(userRegistartionDTO.Email)
	c.JSON(http.StatusOK, gin.H{
		"message": "please check your email for verification",
		"userId":  createdUser.Id,
	})

}
