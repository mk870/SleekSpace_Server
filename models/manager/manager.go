package manager

import baseModel "SleekSpace/models"

type Manager struct {
	baseModel.MyModel
	Id                    int                    `json:"id" gorm:"primary_key"`
	UserId                int                    `json:"userId"`
	Name                  string                 `json:"name" validate:"required,min=2,max=50"`
	Email                 string                 `json:"email"`
	ProfilePicture        ManagerProfilePicture  `gorm:"ForeignKey:ManagerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ManagerContactNumbers []ManagerContactNumber `gorm:"ForeignKey:ManagerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
