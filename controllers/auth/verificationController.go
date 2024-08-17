package controllers

import (
	userService "SleekSpace/services/user"

	"github.com/gin-gonic/gin"
)

func RegistrationVerificationCodeValidation(router *gin.Engine) {
	router.POST("/verification-code/registration", userService.VerifyCodeForRegistration)
}
func SecurityVerificationCodeValidation(router *gin.Engine) {
	router.POST("/verification-code/security", userService.VerifyCodeForSecurity)
}
func ResendVerificationCode(router *gin.Engine) {
	router.GET("/resend-verification-code/:id", userService.ResendVerificationCode)
}

func CreateVerificationCode(router *gin.Engine) {
	router.POST("/verification-code", userService.CreateVerificationCode)
}
