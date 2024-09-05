package commercial

import (
	managerDtos "SleekSpace/dtos/manager"
	insightsDtos "SleekSpace/dtos/property/insights"
	locationDtos "SleekSpace/dtos/property/location"
	imageorvideoDtos "SleekSpace/dtos/property/media"
)

type CommercialForRentPropertyCreationDto struct {
	ManagerId        int                                                                `json:"managerId"`
	NumberOfRooms    int                                                                `json:"numberOfRooms"`
	RentAmount       int                                                                `json:"rentAmount"`
	SizeNumber       int                                                                `json:"sizeNumber"`
	IsFullSpace      bool                                                               `json:"isFullSpace"`
	HasElectricity   bool                                                               `json:"hasElectricity"`
	HasWater         bool                                                               `json:"hasWater"`
	SizeDimensions   string                                                             `json:"sizeDimensions"`
	Status           string                                                             `json:"status"`
	YearBuilt        string                                                             `json:"yearBuilt"`
	Stories          string                                                             `json:"stories"`
	Type             string                                                             `json:"type"`
	InteriorFeatures string                                                             `json:"interiorFeatures"`
	ExteriorFeatures string                                                             `json:"exteriorFeatures"`
	OtherDetails     string                                                             `json:"otherDetails"`
	PropertyLocation locationDtos.PropertyLocationCreationDto                           `json:"propertyLocation"`
	Media            []imageorvideoDtos.PropertyImageOrVideoCreationWithNoPropertyIdDto `json:"media"`
}

type CommercialForRentPropertyResponseDto struct {
	Id               int                                                `json:"id"`
	ManagerId        int                                                `json:"managerId"`
	UniqueId         int                                                `json:"uniqueId"`
	NumberOfRooms    int                                                `json:"numberOfRooms"`
	RentAmount       int                                                `json:"rentAmount"`
	SizeNumber       int                                                `json:"sizeNumber"`
	IsFullSpace      bool                                               `json:"isFullSpace"`
	HasElectricity   bool                                               `json:"hasElectricity"`
	HasWater         bool                                               `json:"hasWater"`
	SizeDimensions   string                                             `json:"sizeDimensions"`
	Status           string                                             `json:"status"`
	YearBuilt        string                                             `json:"yearBuilt"`
	Stories          string                                             `json:"stories"`
	Type             string                                             `json:"type"`
	InteriorFeatures string                                             `json:"interiorFeatures"`
	ExteriorFeatures string                                             `json:"exteriorFeatures"`
	OtherDetails     string                                             `json:"otherDetails"`
	PostedTime       string                                             `json:"postedTime"`
	PropertyLocation locationDtos.PropertyLocationUpdateAndResponseDto  `json:"propertyLocation"`
	Insights         insightsDtos.PropertyInsightsUpdateAndResponseDto  `json:"insights"`
	Media            []imageorvideoDtos.PropertyImageOrVideoResponseDto `json:"media"`
}

type CommercialForRentPropertyWithManagerResponseDto struct {
	Id               int                                                `json:"id"`
	ManagerId        int                                                `json:"managerId"`
	UniqueId         int                                                `json:"uniqueId"`
	NumberOfRooms    int                                                `json:"numberOfRooms"`
	RentAmount       int                                                `json:"rentAmount"`
	SizeNumber       int                                                `json:"sizeNumber"`
	IsFullSpace      bool                                               `json:"isFullSpace"`
	HasElectricity   bool                                               `json:"hasElectricity"`
	HasWater         bool                                               `json:"hasWater"`
	SizeDimensions   string                                             `json:"sizeDimensions"`
	Status           string                                             `json:"status"`
	YearBuilt        string                                             `json:"yearBuilt"`
	Stories          string                                             `json:"stories"`
	Type             string                                             `json:"type"`
	InteriorFeatures string                                             `json:"interiorFeatures"`
	ExteriorFeatures string                                             `json:"exteriorFeatures"`
	OtherDetails     string                                             `json:"otherDetails"`
	PostedTime       string                                             `json:"postedTime"`
	PropertyLocation locationDtos.PropertyLocationUpdateAndResponseDto  `json:"propertyLocation"`
	Insights         insightsDtos.PropertyInsightsUpdateAndResponseDto  `json:"insights"`
	Media            []imageorvideoDtos.PropertyImageOrVideoResponseDto `json:"media"`
	Manager          managerDtos.ManagerResponseDTO                     `json:"manager"`
}

type CommercialForRentPropertyUpdateDto struct {
	Id               int    `json:"id"`
	ManagerId        int    `json:"managerId"`
	UniqueId         int    `json:"uniqueId"`
	NumberOfRooms    int    `json:"numberOfRooms"`
	RentAmount       int    `json:"rentAmount"`
	SizeNumber       int    `json:"sizeNumber"`
	IsFullSpace      bool   `json:"isFullSpace"`
	HasElectricity   bool   `json:"hasElectricity"`
	HasWater         bool   `json:"hasWater"`
	SizeDimensions   string `json:"sizeDimensions"`
	Status           string `json:"status"`
	YearBuilt        string `json:"yearBuilt"`
	Stories          string `json:"stories"`
	Type             string `json:"type"`
	InteriorFeatures string `json:"interiorFeatures"`
	ExteriorFeatures string `json:"exteriorFeatures"`
	OtherDetails     string `json:"otherDetails"`
}
