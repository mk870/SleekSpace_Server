package models

type Manager struct {
	MyModel
	Id                    int                    `json:"id" gorm:"primary_key"`
	UserId                int                    `json:"userId"`
	Name                  string                 `json:"name" validate:"required,min=2,max=50"`
	Email                 string                 `json:"email"`
	Avatar                string                 `json:"avatar"`
	ManagerContactNumbers []ManagerContactNumber `gorm:"ForeignKey:ManagerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CACADE"`
}
