package residential

import (
	insightsDtos "SleekSpace/dtos/property/insights"
	locationDtos "SleekSpace/dtos/property/location"
	imageorvideoDtos "SleekSpace/dtos/property/media"
)

type ResidentialPropertyForRentCreationDto struct {
	ManagerId        int                                                                `json:"managerId"`
	NumberOfRooms    int                                                                `json:"numberOfRooms"`
	RentAmount       int                                                                `json:"rentAmount"`
	NumberOfGarages  int                                                                `json:"numberOfGarages"`
	SizeNumber       int                                                                `json:"sizeNumber"`
	IsFullHouse      bool                                                               `json:"isFullHouse"`
	HasElectricity   bool                                                               `json:"hasElectricity"`
	HasWater         bool                                                               `json:"hasWater"`
	HasSwimmingPool  bool                                                               `json:"hasSwimmingPool"`
	Bedrooms         string                                                             `json:"bedrooms"`
	Bathrooms        string                                                             `json:"bathrooms"`
	Status           string                                                             `json:"status"`
	YearBuilt        string                                                             `json:"yearBuilt"`
	Stories          string                                                             `json:"stories"`
	Type             string                                                             `json:"type"`
	SizeDimensions   string                                                             `json:"sizeDimensions"`
	PropertyLocation locationDtos.PropertyLocationCreationDto                           `json:"propertyLocation"`
	Media            []imageorvideoDtos.PropertyImageOrVideoCreationWithNoPropertyIdDto `json:"media"`
}

type ResidentialPropertyForRentResponseDto struct {
	Id               int                                                         `json:"id"`
	ManagerId        int                                                         `json:"managerId"`
	UniqueId         int                                                         `json:"uniqueId"`
	NumberOfRooms    int                                                         `json:"numberOfRooms"`
	RentAmount       int                                                         `json:"rentAmount"`
	NumberOfGarages  int                                                         `json:"numberOfGarages"`
	SizeNumber       int                                                         `json:"sizeNumber"`
	IsFullHouse      bool                                                        `json:"isFullHouse"`
	HasElectricity   bool                                                        `json:"hasElectricity"`
	HasWater         bool                                                        `json:"hasWater"`
	HasSwimmingPool  bool                                                        `json:"hasSwimmingPool"`
	Bedrooms         string                                                      `json:"bedrooms"`
	Bathrooms        string                                                      `json:"bathrooms"`
	Status           string                                                      `json:"status"`
	YearBuilt        string                                                      `json:"yearBuilt"`
	Stories          string                                                      `json:"stories"`
	Type             string                                                      `json:"type"`
	SizeDimensions   string                                                      `json:"sizeDimensions"`
	PropertyLocation locationDtos.PropertyLocationUpdateAndResponseDto           `json:"propertyLocation"`
	Insights         insightsDtos.PropertyInsightsUpdateAndResponseDto           `json:"insights"`
	Media            []imageorvideoDtos.PropertyImageOrVideoUpdateAndResponseDto `json:"media"`
}

type ResidentialPropertyForRentUpdateDto struct {
	Id              int    `json:"id"`
	ManagerId       int    `json:"managerId"`
	UniqueId        int    `json:"uniqueId"`
	NumberOfRooms   int    `json:"numberOfRooms"`
	RentAmount      int    `json:"rentAmount"`
	NumberOfGarages int    `json:"numberOfGarages"`
	SizeNumber      int    `json:"sizeNumber"`
	IsFullHouse     bool   `json:"isFullHouse"`
	HasElectricity  bool   `json:"hasElectricity"`
	HasWater        bool   `json:"hasWater"`
	HasSwimmingPool bool   `json:"hasSwimmingPool"`
	Bedrooms        string `json:"bedrooms"`
	Bathrooms       string `json:"bathrooms"`
	Status          string `json:"status"`
	YearBuilt       string `json:"yearBuilt"`
	Stories         string `json:"stories"`
	Type            string `json:"type"`
	SizeDimensions  string `json:"sizeDimensions"`
}