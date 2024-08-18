package models

import baseModel "SleekSpace/models"

type RentalProperty struct {
	baseModel.MyModel
	Id               int                    `json:"id" gorm:"primary_key"`
	ManagerId        int                    `json:"managerId"`
	NumberOfRooms    int                    `json:"numberOfRooms"`
	RentAmount       int                    `json:"rentAmount"`
	NumberOfGarages  int                    `json:"numberOfGarages"`
	IsFullHouse      bool                   `json:"isFullHouse"`
	HasElectricity   bool                   `json:"hasElectricity"`
	HasWater         bool                   `json:"hasWater"`
	HasSwimmingPool  bool                   `json:"hasSwimmingPool"`
	Bedrooms         string                 `json:"bedrooms"`
	Bathrooms        string                 `json:"bathrooms"`
	Status           string                 `json:"status"`
	YearBuilt        string                 `json:"yearBuilt"`
	Stories          string                 `json:"stories"`
	Type             string                 `json:"type"`
	Size             string                 `json:"size"`
	Location         PropertyLocation       `gorm:"ForeignKey:PropertyId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyInsights PropertyInsights       `gorm:"ForeignKey:PropertyId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyMedia    []PropertyImageOrVideo `gorm:"ForeignKey:PropertyId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
