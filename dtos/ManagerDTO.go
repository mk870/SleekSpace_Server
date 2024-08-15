package dtos

import managerModels "SleekSpace/models/manager"

type ManagerDTO struct {
	UserId   int                                  `json:"userId"`
	Name     string                               `json:"name" validate:"required,min=2,max=50"`
	Email    string                               `json:"email"`
	Avatar   string                               `json:"avatar"`
	Contacts []managerModels.ManagerContactNumber `json:"contacts"`
}

type ManagerResponseDTO struct {
	Id       int                       `json:"id"`
	UserId   int                       `json:"userId"`
	Name     string                    `json:"name" validate:"required,min=2,max=50"`
	Email    string                    `json:"email"`
	Avatar   string                    `json:"avatar"`
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
