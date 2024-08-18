package models

import baseModel "SleekSpace/models"

type PropertyStand struct {
	baseModel.MyModel
	Id                 int                    `json:"id" gorm:"primary_key"`
	ManagerId          int                    `json:"managerId"`
	Price              int                    `json:"price"`
	AreaHasElectricity bool                   `json:"areaHasElectricity"`
	IsServiced         bool                   `json:"isServiced"`
	IsNegotiable       bool                   `json:"isNegotiable"`
	Status             string                 `json:"status"`
	Level              string                 `json:"level"`
	Size               string                 `json:"size"`
	Type               string                 `json:"type"`
	Location           PropertyLocation       `gorm:"ForeignKey:PropertyId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyInsights   PropertyInsights       `gorm:"ForeignKey:PropertyId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyMedia      []PropertyImageOrVideo `gorm:"ForeignKey:PropertyId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
