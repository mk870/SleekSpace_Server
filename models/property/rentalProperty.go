package models

import baseModel "SleekSpace/models"

type RentalProperty struct {
	baseModel.MyModel
	Id               int              `json:"id" gorm:"primary_key"`
	ManagerId        int              `json:"managerId" gorm:"column:managerId"`
	NumberOfRooms    int              `json:"numberOfRooms"`
	Price            int              `json:"price"`
	Status           string           `json:"status"`
	Images           []string         `json:"images" gorm:"serializer:json"`
	Bedrooms         string           `json:"bedrooms"`
	Bathrooms        string           `json:"bathrooms"`
	NumberOfGarages  int              `json:"numberOfGarages"`
	HasSwimmingPool  bool             `json:"hasSwimmingPool"`
	Size             string           `json:"size"`
	Location         PropertyLocation `gorm:"ForeignKey:PropertyId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyInsights PropertyInsights `gorm:"ForeignKey:PropertyId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
