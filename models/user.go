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
	Id                     int              `json:"id" gorm:"primary_key"`
	GivenName              string           `json:"givenName" validate:"required,min=2,max=50"`
	FamilyName             string           `json:"familyName" validate:"required,min=2,max=50"`
	Email                  string           `json:"email" gorm:"unique" validate:"email,required"`
	Password               string           `json:"password"`
	AccessToken            string           `json:"accessToken"`
	SocialMediaProvider    string           `json:"socialMediaProvider"`
	Avatar                 string           `json:"avatar" gorm:"nullable"`
	IsActive               bool             `json:"isActive"`
	IsSocialsAuthenticated bool             `json:"isSocialsAuthenticated"`
	Location               Location         `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CACADE"`
	ContactNumbers         []ContactNumber  `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CACADE"`
	RegistrationCode       VerificationCode `gorm:"ForeignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CACADE"`
}
