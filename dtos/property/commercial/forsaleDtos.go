package commercial

import (
	managerDtos "SleekSpace/dtos/manager"
	insightsDtos "SleekSpace/dtos/property/insights"
	locationDtos "SleekSpace/dtos/property/location"
	imageorvideoDtos "SleekSpace/dtos/property/media"
)

type CommercialForSalePropertyCreationDto struct {
	ManagerId        int                                                                `json:"managerId"`
	NumberOfRooms    int                                                                `json:"numberOfRooms"`
	Price            int                                                                `json:"price"`
	SizeNumber       int                                                                `json:"sizeNumber"`
	HasElectricity   bool                                                               `json:"hasElectricity"`
	HasWater         bool                                                               `json:"hasWater"`
	IsNegotiable     bool                                                               `json:"isNegotiable"`
	Status           string                                                             `json:"status"`
	YearBuilt        string                                                             `json:"yearBuilt"`
	Stories          string                                                             `json:"stories"`
	Type             string                                                             `json:"type"`
	SizeDimensions   string                                                             `json:"sizeDimensions"`
	InteriorFeatures string                                                             `json:"interiorFeatures"`
	ExteriorFeatures string                                                             `json:"exteriorFeatures"`
	OtherDetails     string                                                             `json:"otherDetails"`
	PropertyLocation locationDtos.PropertyLocationCreationDto                           `json:"propertyLocation"`
	Media            []imageorvideoDtos.PropertyImageOrVideoCreationWithNoPropertyIdDto `json:"media"`
}

type CommercialForSalePropertyResponseDto struct {
	Id               int                                                         `json:"id"`
	ManagerId        int                                                         `json:"managerId"`
	UniqueId         int                                                         `json:"uniqueId"`
	NumberOfRooms    int                                                         `json:"numberOfRooms"`
	Price            int                                                         `json:"price"`
	SizeNumber       int                                                         `json:"sizeNumber"`
	HasElectricity   bool                                                        `json:"hasElectricity"`
	IsNegotiable     bool                                                        `json:"isNegotiable"`
	HasWater         bool                                                        `json:"hasWater"`
	Status           string                                                      `json:"status"`
	YearBuilt        string                                                      `json:"yearBuilt"`
	Stories          string                                                      `json:"stories"`
	Type             string                                                      `json:"type"`
	SizeDimensions   string                                                      `json:"sizeDimensions"`
	InteriorFeatures string                                                      `json:"interiorFeatures"`
	ExteriorFeatures string                                                      `json:"exteriorFeatures"`
	OtherDetails     string                                                      `json:"otherDetails"`
	PostedTime       string                                                      `json:"postedTime"`
	PropertyLocation locationDtos.PropertyLocationUpdateAndResponseDto           `json:"propertyLocation"`
	Insights         insightsDtos.PropertyInsightsUpdateAndResponseDto           `json:"insights"`
	Media            []imageorvideoDtos.PropertyImageOrVideoUpdateAndResponseDto `json:"media"`
}

type CommercialForSalePropertyWithManagerResponseDto struct {
	Id               int                                                         `json:"id"`
	ManagerId        int                                                         `json:"managerId"`
	UniqueId         int                                                         `json:"uniqueId"`
	NumberOfRooms    int                                                         `json:"numberOfRooms"`
	Price            int                                                         `json:"price"`
	SizeNumber       int                                                         `json:"sizeNumber"`
	HasElectricity   bool                                                        `json:"hasElectricity"`
	IsNegotiable     bool                                                        `json:"isNegotiable"`
	HasWater         bool                                                        `json:"hasWater"`
	Status           string                                                      `json:"status"`
	YearBuilt        string                                                      `json:"yearBuilt"`
	Stories          string                                                      `json:"stories"`
	Type             string                                                      `json:"type"`
	SizeDimensions   string                                                      `json:"sizeDimensions"`
	InteriorFeatures string                                                      `json:"interiorFeatures"`
	ExteriorFeatures string                                                      `json:"exteriorFeatures"`
	OtherDetails     string                                                      `json:"otherDetails"`
	PostedTime       string                                                      `json:"postedTime"`
	PropertyLocation locationDtos.PropertyLocationUpdateAndResponseDto           `json:"propertyLocation"`
	Insights         insightsDtos.PropertyInsightsUpdateAndResponseDto           `json:"insights"`
	Media            []imageorvideoDtos.PropertyImageOrVideoUpdateAndResponseDto `json:"media"`
	Manager          managerDtos.ManagerResponseDTO                              `json:"manager"`
}

type CommercialForSalePropertyUpdateDto struct {
	Id               int    `json:"id"`
	ManagerId        int    `json:"managerId"`
	UniqueId         int    `json:"uniqueId"`
	NumberOfRooms    int    `json:"numberOfRooms"`
	Price            int    `json:"price"`
	SizeNumber       int    `json:"sizeNumber"`
	HasElectricity   bool   `json:"hasElectricity"`
	HasWater         bool   `json:"hasWater"`
	IsNegotiable     bool   `json:"isNegotiable"`
	Status           string `json:"status"`
	YearBuilt        string `json:"yearBuilt"`
	Stories          string `json:"stories"`
	Type             string `json:"type"`
	SizeDimensions   string `json:"sizeDimensions"`
	InteriorFeatures string `json:"interiorFeatures"`
	ExteriorFeatures string `json:"exteriorFeatures"`
	OtherDetails     string `json:"otherDetails"`
}
