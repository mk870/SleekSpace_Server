package models

import baseModel "SleekSpace/models"

type PropertyStand struct {
	baseModel.MyModel
	Id                 int                    `json:"id" gorm:"primary_key"`
	ManagerId          int                    `json:"managerId"`
	UniqueId           int                    `json:"uniqueId"`
	Price              int                    `json:"price"`
	SizeNumber         int                    `json:"sizeNumber"`
	AreaHasElectricity bool                   `json:"areaHasElectricity"`
	IsServiced         bool                   `json:"isServiced"`
	IsNegotiable       bool                   `json:"isNegotiable"`
	Status             string                 `json:"status"`
	Level              string                 `json:"level"`
	SizeDimensions     string                 `json:"sizeDimensions"`
	Type               string                 `json:"type"`
	Location           PropertyLocation       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyInsights   PropertyInsights       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyMedia      []PropertyImageOrVideo `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
