package manager

import (
	baseModel "SleekSpace/models"
	propertyModel "SleekSpace/models/property"
)

type Manager struct {
	baseModel.MyModel
	Id                         int                                        `json:"id" gorm:"primary_key"`
	UserId                     int                                        `json:"userId"`
	Name                       string                                     `json:"name" validate:"required,min=2,max=50"`
	Email                      string                                     `json:"email"`
	ProfilePicture             ManagerProfilePicture                      `gorm:"ForeignKey:ManagerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ManagerContactNumbers      []ManagerContactNumber                     `gorm:"ForeignKey:ManagerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CommercialForSaleProperty  []propertyModel.CommercialForSaleProperty  `gorm:"ForeignKey:ManagerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CommercialRentalProperty   []propertyModel.CommercialRentalProperty   `gorm:"ForeignKey:ManagerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ResidentialPropertyForSale []propertyModel.ResidentialPropertyForSale `gorm:"ForeignKey:ManagerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ResidentialRentalProperty  []propertyModel.ResidentialRentalProperty  `gorm:"ForeignKey:ManagerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyStand              []propertyModel.PropertyStand              `gorm:"ForeignKey:ManagerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	LandForSaleProperty        []propertyModel.LandForSaleProperty        `gorm:"ForeignKey:ManagerId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
