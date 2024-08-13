package repositories

import (
	"SleekSpace/db"
	"SleekSpace/models"
)

func CreateVerificationCode(user *models.User, verificationCode models.VerificationCode) bool {
	db.DB.Model(user).Association("RegistrationCode").Replace(verificationCode)
	return true
}

func GetVerificationCodeByUserId(userId string) models.VerificationCode {
	var user models.User
	err := db.DB.Preload("RegistrationCode").First(&user, userId)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return user.RegistrationCode
}

func GetVerificationCodeById(id string) models.VerificationCode {
	var code models.VerificationCode
	err := db.DB.First(&code, id)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return code
}

func UpdateVerificationCode(verificationCodeUpdate *models.VerificationCode) bool {
	db.DB.Save(verificationCodeUpdate)
	return true
}

func AllVerificationCodes() []models.VerificationCode {
	var codes []models.VerificationCode
	err := db.DB.Find(&codes)
	if err != nil {
		println(err.Name(), err.Statement)
	}
	return codes
}

func DeleteVerficationCode(userId int) bool {
	db.DB.Unscoped().Delete(&models.VerificationCode{}, userId)
	return true
}
