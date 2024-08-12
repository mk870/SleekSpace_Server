package controllers

import (
	"SleekSpace/services"

	"github.com/gin-gonic/gin"
)

func RegistrationVerificationCodeValidation(router *gin.Engine) {
	router.POST("/verification-code/registration", services.VerifyCodeForRegistration)
}
func SecurityVerificationCodeValidation(router *gin.Engine) {
	router.POST("/verification-code/security", services.VerifyCodeForSecurity)
}
func ResendVerificationCode(router *gin.Engine) {
	router.GET("/resend-verification-code/:id", services.ResendVerificationCode)
}

func CreateVerificationCode(router *gin.Engine) {
	router.POST("/verification-code", services.CreateVerificationCode)
}

func AllVerificationCodes(router *gin.Engine) {
	router.GET("/verification-codes", services.GetAllCodes)
}

func DeleteVerificationCode(router *gin.Engine) {
	router.DELETE("/verification-code/:id", services.DeleteVerificationCode)
}
