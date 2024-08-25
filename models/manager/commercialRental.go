package manager

import (
	baseModel "SleekSpace/models"
	propertyModels "SleekSpace/models/property"
)

type CommercialRentalProperty struct {
	baseModel.MyModel
	Id               int                                   `json:"id" gorm:"primary_key"`
	ManagerId        int                                   `json:"managerId"`
	UniqueId         int                                   `json:"uniqueId"`
	NumberOfRooms    int                                   `json:"numberOfRooms"`
	RentAmount       int                                   `json:"rentAmount"`
	SizeNumber       int                                   `json:"sizeNumber"`
	IsFullSpace      bool                                  `json:"isFullSpace"`
	HasElectricity   bool                                  `json:"hasElectricity"`
	HasWater         bool                                  `json:"hasWater"`
	SizeDimensions   string                                `json:"sizeDimensions"`
	Status           string                                `json:"status"`
	YearBuilt        string                                `json:"yearBuilt"`
	Stories          string                                `json:"stories"`
	Type             string                                `json:"type"`
	InteriorFeatures string                                `gorm:"type:text"`
	ExteriorFeatures string                                `gorm:"type:text"`
	OtherDetails     string                                `gorm:"type:text"`
	Manager          Manager                               `json:"manager"`
	Location         propertyModels.PropertyLocation       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyInsights propertyModels.PropertyInsights       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyMedia    []propertyModels.PropertyImageOrVideo `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
