package dtos

import "SleekSpace/models"

type ManagerDTO struct {
	UserId   int                           `json:"userId"`
	Name     string                        `json:"name" validate:"required,min=2,max=50"`
	Email    string                        `json:"email"`
	Contacts []models.ManagerContactNumber `json:"contacts"`
}

type ManagerResponseDTO struct {
	Id       int                       `json:"id"`
	UserId   int                       `json:"userId"`
	Name     string                    `json:"name" validate:"required,min=2,max=50"`
	Email    string                    `json:"email"`
	Contacts []ManagerContactNumberDTO `json:"contacts"`
}

type ManagerContactNumberDTO struct {
	Id           int    `json:"id"`
	ManagerId    int    `json:"managerId"`
	Type         string `json:"type"`
	Number       string `json:"number"`
	CountryCode  string `json:"countryCode"`
	CountryAbbrv string `json:"countryAbbrv"`
}
