package stand

import (
	imageorvideoDtos "SleekSpace/dtos/property/imageorvideo"
	insightsDtos "SleekSpace/dtos/property/insights"
	locationDtos "SleekSpace/dtos/property/location"
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
	PropertyLocation   locationDtos.PropertyLocationCreationDto                           `json:"propertyLocation"`
	Media              []imageorvideoDtos.PropertyImageOrVideoCreationWithNoPropertyIdDto `json:"media"`
}

type StandResponseDTO struct {
	Id                 int                                                         `json:"id"`
	ManagerId          int                                                         `json:"managerId"`
	Price              int                                                         `json:"price"`
	SizeNumber         int                                                         `json:"sizeNumber"`
	AreaHasElectricity bool                                                        `json:"areaHasElectricity"`
	IsServiced         bool                                                        `json:"isServiced"`
	IsNegotiable       bool                                                        `json:"isNegotiable"`
	Status             string                                                      `json:"status"`
	Level              string                                                      `json:"level"`
	SizeDimensions     string                                                      `json:"sizeDimensions"`
	Type               string                                                      `json:"type"`
	PropertyLocation   locationDtos.PropertyLocationUpdateAndResponseDto           `json:"propertyLocation"`
	Insights           insightsDtos.PropertyInsightsUpdateAndResponseDto           `json:"insights"`
	Media              []imageorvideoDtos.PropertyImageOrVideoUpdateAndResponseDto `json:"media"`
}

type StandUpdateDTO struct {
	Id                 int    `json:"id"`
	ManagerId          int    `json:"managerId"`
	Price              int    `json:"price"`
	SizeNumber         int    `json:"sizeNumber"`
	AreaHasElectricity bool   `json:"areaHasElectricity"`
	IsServiced         bool   `json:"isServiced"`
	IsNegotiable       bool   `json:"isNegotiable"`
	Status             string `json:"status"`
	Level              string `json:"level"`
	SizeDimensions     string `json:"sizeDimensions"`
	Type               string `json:"type"`
}
