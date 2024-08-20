package land

import (
	imageorvideoDtos "SleekSpace/dtos/property/imageorvideo"
	insightsDtos "SleekSpace/dtos/property/insights"
	locationDtos "SleekSpace/dtos/property/location"
)

type LandForSalePropertyCreationDto struct {
	ManagerId          int                                                                `json:"managerId"`
	Price              int                                                                `json:"price"`
	SizeNumber         int                                                                `json:"sizeNumber"`
	AreaHasElectricity bool                                                               `json:"areaHasElectricity"`
	HasWater           bool                                                               `json:"hasWater"`
	Status             string                                                             `json:"status"`
	Type               string                                                             `json:"type"`
	SizeDimensions     string                                                             `json:"sizeDimensions"`
	PropertyLocation   locationDtos.PropertyLocationCreationDto                           `json:"propertyLocation"`
	Media              []imageorvideoDtos.PropertyImageOrVideoCreationWithNoPropertyIdDto `json:"media"`
}

type LandForSalePropertyResponseDto struct {
	Id                 int                                                         `json:"id"`
	ManagerId          int                                                         `json:"managerId"`
	UniqueId           int                                                         `json:"uniqueId"`
	Price              int                                                         `json:"price"`
	SizeNumber         int                                                         `json:"sizeNumber"`
	AreaHasElectricity bool                                                        `json:"areaHasElectricity"`
	HasWater           bool                                                        `json:"hasWater"`
	Status             string                                                      `json:"status"`
	Type               string                                                      `json:"type"`
	SizeDimensions     string                                                      `json:"sizeDimensions"`
	PropertyLocation   locationDtos.PropertyLocationUpdateAndResponseDto           `json:"propertyLocation"`
	Insights           insightsDtos.PropertyInsightsUpdateAndResponseDto           `json:"insights"`
	Media              []imageorvideoDtos.PropertyImageOrVideoUpdateAndResponseDto `json:"media"`
}

type LandForSalePropertyUpdateDto struct {
	Id                 int    `json:"id"`
	ManagerId          int    `json:"managerId"`
	UniqueId           int    `json:"uniqueId"`
	Price              int    `json:"price"`
	SizeNumber         int    `json:"sizeNumber"`
	AreaHasElectricity bool   `json:"areaHasElectricity"`
	HasWater           bool   `json:"hasWater"`
	Status             string `json:"status"`
	Type               string `json:"type"`
	SizeDimensions     string `json:"sizeDimensions"`
}
