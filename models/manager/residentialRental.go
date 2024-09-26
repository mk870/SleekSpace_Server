package manager

import (
	baseModel "SleekSpace/models"
	propertyModels "SleekSpace/models/property"
)

type ResidentialRentalProperty struct {
	baseModel.MyModel
	Id                     int                                   `json:"id" gorm:"primary_key"`
	ManagerId              int                                   `json:"managerId"`
	UniqueId               int                                   `json:"uniqueId"`
	NumberOfRoomsToLet     int                                   `json:"numberOfRoomsToLet"`
	TotalNumberOfRooms     int                                   `json:"totalNumberOfRooms"`
	RentAmount             int                                   `json:"rentAmount"`
	NumberOfGarages        int                                   `json:"numberOfGarages"`
	SizeNumber             int                                   `json:"sizeNumber"`
	Bedrooms               int                                   `json:"bedrooms"`
	Bathrooms              int                                   `json:"bathrooms"`
	Stories                int                                   `json:"stories"`
	YearBuilt              int                                   `json:"yearBuilt"`
	IsFullHouse            bool                                  `json:"isFullHouse"`
	HasElectricity         bool                                  `json:"hasElectricity"`
	HasWater               bool                                  `json:"hasWater"`
	IsPaved                bool                                  `json:"isPaved"`
	HasBoreHole            bool                                  `json:"hasBoreHole"`
	IsPlustered            bool                                  `json:"isPlustered"`
	IsPainted              bool                                  `json:"isPainted"`
	IsTiled                bool                                  `json:"isTiled"`
	HasCeiling             bool                                  `json:"hasCeiling"`
	HasSwimmingPool        bool                                  `json:"hasSwimmingPool"`
	TypeOfExteriorSecurity string                                `json:"typeOfExteriorSecurity"`
	TenantRequirements     string                                `json:"tenantRequirements"`
	Status                 string                                `json:"status"`
	Type                   string                                `json:"type"`
	MarketingStatement     string                                `gorm:"type:text"`
	SizeDimensions         string                                `json:"sizeDimensions"`
	Currency               string                                `json:"currency"`
	OtherInteriorFeatures  string                                `gorm:"type:text"`
	OtherExteriorFeatures  string                                `gorm:"type:text"`
	Manager                Manager                               `json:"manager"`
	Location               propertyModels.PropertyLocation       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyInsights       propertyModels.PropertyInsights       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyMedia          []propertyModels.PropertyImageOrVideo `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PropertyReport         []propertyModels.PropertyReport       `gorm:"ForeignKey:PropertyId;references:UniqueId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
