package manager

import (
	baseModel "SleekSpace/models"
	propertyModels "SleekSpace/models/property"
)

type ResidentialPropertyForSale struct {
	baseModel.MyModel
	Id               int                                   `json:"id" gorm:"primary_key"`
	ManagerId        int                                   `json:"managerId"`
	UniqueId         int                                   `json:"uniqueId"`
	NumberOfRooms    int                                   `json:"numberOfRooms"`
	Price            int                                   `json:"price"`
	NumberOfGarages  int                                   `json:"numberOfGarages"`
	SizeNumber       int                                   `json:"sizeNumber"`
	HasSwimmingPool  bool                                  `json:"hasSwimmingPool"`
	HasElectricity   bool                                  `json:"hasElectricity"`
	HasWater         bool                                  `json:"hasWater"`
	IsNegotiable     bool                                  `json:"isNegotiable"`
	Status           string                                `json:"status"`
	Bedrooms         string                                `json:"bedrooms"`
	Bathrooms        string                                `json:"bathrooms"`
	YearBuilt        string                                `json:"yearBuilt"`
	Stories          string                                `json:"stories"`
	Type             string                                `json:"type"`
	SizeDimensions   string                                `json:"sizeDimensions"`
	InteriorFeatures string                                `gorm:"type:text"`
	ExteriorFeatures string                                `gorm:"type:text"`
	OtherDetails     string                                `gorm:"type:text"`
	Manager          Manager                               `json:"manager"`
	Location         propertyModels.PropertyLocation       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyInsights propertyModels.PropertyInsights       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyMedia    []propertyModels.PropertyImageOrVideo `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
