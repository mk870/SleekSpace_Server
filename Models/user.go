package models

import (
	"time"

	"gorm.io/gorm"
)

type MyModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	MyModel
	Id               int              `json:"id" gorm:"primary_key"`
	GivenName        string           `json:"givenName" validate:"required,min=2,max=50"`
	FamilyName       string           `json:"familyName" validate:"required,min=2,max=50"`
	Email            string           `json:"email" gorm:"unique" validate:"email,required"`
	Password         string           `json:"password"`
	Photo            string           `json:"photo" gorm:"nullable"`
	IsActive         bool             `json:"isActive"`
	RegistrationCode VerificationCode `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
