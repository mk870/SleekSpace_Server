package models

import baseModel "SleekSpace/models"

type CommercialForSaleProperty struct {
	baseModel.MyModel
	Id               int                    `json:"id" gorm:"primary_key"`
	ManagerId        int                    `json:"managerId"`
	NumberOfRooms    int                    `json:"numberOfRooms"`
	Price            int                    `json:"price"`
	SizeNumber       int                    `json:"sizeNumber"`
	HasElectricity   bool                   `json:"hasElectricity"`
	HasWater         bool                   `json:"hasWater"`
	Status           string                 `json:"status"`
	YearBuilt        string                 `json:"yearBuilt"`
	Stories          string                 `json:"stories"`
	Type             string                 `json:"type"`
	SizeDimensions   string                 `json:"sizeDimensions"`
	Location         PropertyLocation       `gorm:"ForeignKey:PropertyId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyInsights PropertyInsights       `gorm:"ForeignKey:PropertyId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyMedia    []PropertyImageOrVideo `gorm:"ForeignKey:PropertyId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
