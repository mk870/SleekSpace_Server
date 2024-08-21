package residential

import (
	insightsDtos "SleekSpace/dtos/property/insights"
	locationDtos "SleekSpace/dtos/property/location"
	imageorvideoDtos "SleekSpace/dtos/property/media"
)

type ResidentialPropertyForSaleCreationDto struct {
	ManagerId        int                                                                `json:"managerId"`
	NumberOfRooms    int                                                                `json:"numberOfRooms"`
	Price            int                                                                `json:"price"`
	NumberOfGarages  int                                                                `json:"numberOfGarages"`
	SizeNumber       int                                                                `json:"sizeNumber"`
	HasSwimmingPool  bool                                                               `json:"hasSwimmingPool"`
	HasElectricity   bool                                                               `json:"hasElectricity"`
	HasWater         bool                                                               `json:"hasWater"`
	IsNegotiable     bool                                                               `json:"isNegotiable"`
	Status           string                                                             `json:"status"`
	Bedrooms         string                                                             `json:"bedrooms"`
	Bathrooms        string                                                             `json:"bathrooms"`
	YearBuilt        string                                                             `json:"yearBuilt"`
	Stories          string                                                             `json:"stories"`
	Type             string                                                             `json:"type"`
	SizeDimensions   string                                                             `json:"sizeDimensions"`
	PropertyLocation locationDtos.PropertyLocationCreationDto                           `json:"propertyLocation"`
	Media            []imageorvideoDtos.PropertyImageOrVideoCreationWithNoPropertyIdDto `json:"media"`
}

type ResidentialPropertyForSaleResponseDto struct {
	Id               int                                                         `json:"id"`
	ManagerId        int                                                         `json:"managerId"`
	UniqueId         int                                                         `json:"uniqueId"`
	NumberOfRooms    int                                                         `json:"numberOfRooms"`
	Price            int                                                         `json:"price"`
	NumberOfGarages  int                                                         `json:"numberOfGarages"`
	SizeNumber       int                                                         `json:"sizeNumber"`
	HasSwimmingPool  bool                                                        `json:"hasSwimmingPool"`
	HasElectricity   bool                                                        `json:"hasElectricity"`
	HasWater         bool                                                        `json:"hasWater"`
	IsNegotiable     bool                                                        `json:"isNegotiable"`
	Status           string                                                      `json:"status"`
	Bedrooms         string                                                      `json:"bedrooms"`
	Bathrooms        string                                                      `json:"bathrooms"`
	YearBuilt        string                                                      `json:"yearBuilt"`
	Stories          string                                                      `json:"stories"`
	Type             string                                                      `json:"type"`
	SizeDimensions   string                                                      `json:"sizeDimensions"`
	PropertyLocation locationDtos.PropertyLocationUpdateAndResponseDto           `json:"propertyLocation"`
	Insights         insightsDtos.PropertyInsightsUpdateAndResponseDto           `json:"insights"`
	Media            []imageorvideoDtos.PropertyImageOrVideoUpdateAndResponseDto `json:"media"`
}

type ResidentialPropertyForSaleUpdateDto struct {
	Id              int    `json:"id"`
	ManagerId       int    `json:"managerId"`
	UniqueId        int    `json:"uniqueId"`
	NumberOfRooms   int    `json:"numberOfRooms"`
	Price           int    `json:"price"`
	NumberOfGarages int    `json:"numberOfGarages"`
	SizeNumber      int    `json:"sizeNumber"`
	HasSwimmingPool bool   `json:"hasSwimmingPool"`
	HasElectricity  bool   `json:"hasElectricity"`
	HasWater        bool   `json:"hasWater"`
	IsNegotiable    bool   `json:"isNegotiable"`
	Status          string `json:"status"`
	Bedrooms        string `json:"bedrooms"`
	Bathrooms       string `json:"bathrooms"`
	YearBuilt       string `json:"yearBuilt"`
	Stories         string `json:"stories"`
	Type            string `json:"type"`
	SizeDimensions  string `json:"sizeDimensions"`
}
