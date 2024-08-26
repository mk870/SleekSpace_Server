package stand

import (
	managerDtos "SleekSpace/dtos/manager"
	insightsDtos "SleekSpace/dtos/property/insights"
	locationDtos "SleekSpace/dtos/property/location"
	imageorvideoDtos "SleekSpace/dtos/property/media"
)

type StandCreationDTO struct {
	ManagerId          int                                                                `json:"managerId"`
	Price              int                                                                `json:"price"`
	SizeNumber         int                                                                `json:"sizeNumber"`
	AreaHasElectricity bool                                                               `json:"areaHasElectricity"`
	IsServiced         bool                                                               `json:"isServiced"`
	IsNegotiable       bool                                                               `json:"isNegotiable"`
	Status             string                                                             `json:"status"`
	Level              string                                                             `json:"level"`
	SizeDimensions     string                                                             `json:"sizeDimensions"`
	Type               string                                                             `json:"type"`
	OtherDetails       string                                                             `json:"otherDetails"`
	PropertyLocation   locationDtos.PropertyLocationCreationDto                           `json:"propertyLocation"`
	Media              []imageorvideoDtos.PropertyImageOrVideoCreationWithNoPropertyIdDto `json:"media"`
}

type StandResponseDTO struct {
	Id                 int                                                         `json:"id"`
	ManagerId          int                                                         `json:"managerId"`
	UniqueId           int                                                         `json:"uniqueId"`
	Price              int                                                         `json:"price"`
	SizeNumber         int                                                         `json:"sizeNumber"`
	AreaHasElectricity bool                                                        `json:"areaHasElectricity"`
	IsServiced         bool                                                        `json:"isServiced"`
	IsNegotiable       bool                                                        `json:"isNegotiable"`
	Status             string                                                      `json:"status"`
	Level              string                                                      `json:"level"`
	SizeDimensions     string                                                      `json:"sizeDimensions"`
	Type               string                                                      `json:"type"`
	OtherDetails       string                                                      `json:"otherDetails"`
	PostedTime         string                                                      `json:"postedTime"`
	PropertyLocation   locationDtos.PropertyLocationUpdateAndResponseDto           `json:"propertyLocation"`
	Insights           insightsDtos.PropertyInsightsUpdateAndResponseDto           `json:"insights"`
	Media              []imageorvideoDtos.PropertyImageOrVideoUpdateAndResponseDto `json:"media"`
}

type StandWithManagerResponseDTO struct {
	Id                 int                                                         `json:"id"`
	ManagerId          int                                                         `json:"managerId"`
	UniqueId           int                                                         `json:"uniqueId"`
	Price              int                                                         `json:"price"`
	SizeNumber         int                                                         `json:"sizeNumber"`
	AreaHasElectricity bool                                                        `json:"areaHasElectricity"`
	IsServiced         bool                                                        `json:"isServiced"`
	IsNegotiable       bool                                                        `json:"isNegotiable"`
	Status             string                                                      `json:"status"`
	Level              string                                                      `json:"level"`
	SizeDimensions     string                                                      `json:"sizeDimensions"`
	Type               string                                                      `json:"type"`
	OtherDetails       string                                                      `json:"otherDetails"`
	PostedTime         string                                                      `json:"postedTime"`
	PropertyLocation   locationDtos.PropertyLocationUpdateAndResponseDto           `json:"propertyLocation"`
	Insights           insightsDtos.PropertyInsightsUpdateAndResponseDto           `json:"insights"`
	Media              []imageorvideoDtos.PropertyImageOrVideoUpdateAndResponseDto `json:"media"`
	Manager            managerDtos.ManagerResponseDTO                              `json:"manager"`
}

type StandUpdateDTO struct {
	Id                 int    `json:"id"`
	ManagerId          int    `json:"managerId"`
	UniqueId           int    `json:"uniqueId"`
	Price              int    `json:"price"`
	SizeNumber         int    `json:"sizeNumber"`
	AreaHasElectricity bool   `json:"areaHasElectricity"`
	IsServiced         bool   `json:"isServiced"`
	IsNegotiable       bool   `json:"isNegotiable"`
	Status             string `json:"status"`
	Level              string `json:"level"`
	SizeDimensions     string `json:"sizeDimensions"`
	Type               string `json:"type"`
	OtherDetails       string `json:"otherDetails"`
}
