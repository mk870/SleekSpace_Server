package models

import baseModel "SleekSpace/models"

type LandForSaleProperty struct {
	baseModel.MyModel
	Id                 int                    `json:"id" gorm:"primary_key"`
	ManagerId          int                    `json:"managerId"`
	UniqueId           int                    `json:"uniqueId"`
	Price              int                    `json:"price"`
	SizeNumber         int                    `json:"sizeNumber"`
	AreaHasElectricity bool                   `json:"areaHasElectricity"`
	HasWater           bool                   `json:"hasWater"`
	Status             string                 `json:"status"`
	Type               string                 `json:"type"`
	SizeDimensions     string                 `json:"sizeDimensions"`
	Location           PropertyLocation       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyInsights   PropertyInsights       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyMedia      []PropertyImageOrVideo `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
