package db

import (
	"SleekSpace/models"
	"SleekSpace/utilities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	databaseDetails := utilities.GetEnvVariables().DatabaseDetails
	db, err := gorm.Open(postgres.Open(databaseDetails), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.ManagerContactNumber{})
	db.AutoMigrate(&models.VerificationCode{}, &models.ContactNumber{}, &models.Location{}, &models.Manager{})
	db.AutoMigrate(&models.User{})
	DB = db
}
