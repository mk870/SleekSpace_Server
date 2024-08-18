package models

import baseModel "SleekSpace/models"

type ResidentialRentalProperty struct {
	baseModel.MyModel
	Id               int                    `json:"id" gorm:"primary_key"`
	ManagerId        int                    `json:"managerId"`
	NumberOfRooms    int                    `json:"numberOfRooms"`
	RentAmount       int                    `json:"rentAmount"`
	NumberOfGarages  int                    `json:"numberOfGarages"`
	SizeNumber       int                    `json:"sizeNumber"`
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
	SizeDimensions   string                 `json:"sizeDimensions"`
	Location         PropertyLocation       `gorm:"ForeignKey:PropertyId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyInsights PropertyInsights       `gorm:"ForeignKey:PropertyId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyMedia    []PropertyImageOrVideo `gorm:"ForeignKey:PropertyId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
