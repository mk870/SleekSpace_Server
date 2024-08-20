package models

import baseModel "SleekSpace/models"

type CommercialRentalProperty struct {
	baseModel.MyModel
	Id               int                    `json:"id" gorm:"primary_key"`
	ManagerId        int                    `json:"managerId"`
	UniqueId         int                    `json:"uniqueId"`
	NumberOfRooms    int                    `json:"numberOfRooms"`
	RentAmount       int                    `json:"rentAmount"`
	SizeNumber       int                    `json:"sizeNumber"`
	IsFullSpace      bool                   `json:"isFullSpace"`
	HasElectricity   bool                   `json:"hasElectricity"`
	HasWater         bool                   `json:"hasWater"`
	SizeDimensions   string                 `json:"sizeDimensions"`
	Status           string                 `json:"status"`
	YearBuilt        string                 `json:"yearBuilt"`
	Stories          string                 `json:"stories"`
	Type             string                 `json:"type"`
	Location         PropertyLocation       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyInsights PropertyInsights       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyMedia    []PropertyImageOrVideo `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
