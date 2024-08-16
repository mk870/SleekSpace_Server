package dtos

import managerModels "SleekSpace/models/manager"

type ManagerCreationDTO struct {
	UserId   int                                  `json:"userId"`
	Name     string                               `json:"name" validate:"required,min=2,max=50"`
	Email    string                               `json:"email"`
	Contacts []managerModels.ManagerContactNumber `json:"contacts"`
}

type ManagerResponseDTO struct {
	Id             int                                       `json:"id"`
	UserId         int                                       `json:"userId"`
	Name           string                                    `json:"name" validate:"required,min=2,max=50"`
	Email          string                                    `json:"email"`
	ProfilePicture ManagerProfilePictureResponseAndUpdateDTO `json:"profilePicture"`
	Contacts       []ManagerContactNumberDTO                 `json:"contacts"`
}

type UpdateManagerNameAndEmailDTO struct {
	Name  string `json:"name" validate:"required,min=2,max=50"`
	Email string `json:"email"`
}
