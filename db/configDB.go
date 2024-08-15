package db

import (
	managerModels "SleekSpace/models/manager"
	userModels "SleekSpace/models/user"
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
	db.AutoMigrate(&managerModels.ManagerContactNumber{})
	db.AutoMigrate(&userModels.VerificationCode{}, &userModels.ContactNumber{}, &userModels.Location{}, &managerModels.Manager{})
	db.AutoMigrate(&userModels.User{})
	DB = db
}
