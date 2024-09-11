package property

import (
	commercialPropertyDtos "SleekSpace/dtos/property/commercial"
	managerModels "SleekSpace/models/manager"
	generalUtilities "SleekSpace/utilities/funcs/general"
	managerUtilities "SleekSpace/utilities/funcs/manager"
)

func CommercialPropertyForRentResponse(commercialPropertyForRent managerModels.CommercialRentalProperty) commercialPropertyDtos.CommercialForRentPropertyResponseDto {
	return commercialPropertyDtos.CommercialForRentPropertyResponseDto{
		Id:               commercialPropertyForRent.Id,
		ManagerId:        commercialPropertyForRent.ManagerId,
		UniqueId:         commercialPropertyForRent.UniqueId,
		RentAmount:       commercialPropertyForRent.RentAmount,
		SizeNumber:       commercialPropertyForRent.SizeNumber,
		SizeDimensions:   commercialPropertyForRent.SizeDimensions,
		Status:           commercialPropertyForRent.Status,
		Type:             commercialPropertyForRent.Type,
		YearBuilt:        commercialPropertyForRent.YearBuilt,
		Stories:          commercialPropertyForRent.Stories,
		HasElectricity:   commercialPropertyForRent.HasElectricity,
		HasWater:         commercialPropertyForRent.HasWater,
		NumberOfRooms:    commercialPropertyForRent.NumberOfRooms,
		IsFullSpace:      commercialPropertyForRent.IsFullSpace,
		OtherDetails:     commercialPropertyForRent.OtherDetails,
		ExteriorFeatures: commercialPropertyForRent.ExteriorFeatures,
		InteriorFeatures: commercialPropertyForRent.InteriorFeatures,
		PostedTime:       generalUtilities.GetTimePassed(commercialPropertyForRent.CreatedAt),
		PropertyLocation: PropertyLocationResponse(commercialPropertyForRent.Location),
		Insights:         PropertyInsightsResponse(commercialPropertyForRent.PropertyInsights),
		Media:            ProcessedPropertyImageAndVideosListToResponse(commercialPropertyForRent.PropertyMedia),
	}
}

func CommercialPropertyForRentWithManagerResponse(commercialPropertyForRent managerModels.CommercialRentalProperty) commercialPropertyDtos.CommercialForRentPropertyWithManagerResponseDto {
	return commercialPropertyDtos.CommercialForRentPropertyWithManagerResponseDto{
		Id:               commercialPropertyForRent.Id,
		ManagerId:        commercialPropertyForRent.ManagerId,
		UniqueId:         commercialPropertyForRent.UniqueId,
		RentAmount:       commercialPropertyForRent.RentAmount,
		SizeNumber:       commercialPropertyForRent.SizeNumber,
		SizeDimensions:   commercialPropertyForRent.SizeDimensions,
		Status:           commercialPropertyForRent.Status,
		Type:             commercialPropertyForRent.Type,
		YearBuilt:        commercialPropertyForRent.YearBuilt,
		Stories:          commercialPropertyForRent.Stories,
		HasElectricity:   commercialPropertyForRent.HasElectricity,
		HasWater:         commercialPropertyForRent.HasWater,
		NumberOfRooms:    commercialPropertyForRent.NumberOfRooms,
		IsFullSpace:      commercialPropertyForRent.IsFullSpace,
		OtherDetails:     commercialPropertyForRent.OtherDetails,
		ExteriorFeatures: commercialPropertyForRent.ExteriorFeatures,
		InteriorFeatures: commercialPropertyForRent.InteriorFeatures,
		PostedTime:       generalUtilities.GetTimePassed(commercialPropertyForRent.CreatedAt),
		PropertyLocation: PropertyLocationResponse(commercialPropertyForRent.Location),
		Insights:         PropertyInsightsResponse(commercialPropertyForRent.PropertyInsights),
		Media:            ProcessedPropertyImageAndVideosListToResponse(commercialPropertyForRent.PropertyMedia),
		Manager:          managerUtilities.ManagerResponse(&commercialPropertyForRent.Manager),
	}
}

func CommercialPropertyForSaleWithManagerResponse(commercialPropertyForSale managerModels.CommercialForSaleProperty) commercialPropertyDtos.CommercialForSalePropertyWithManagerResponseDto {
	return commercialPropertyDtos.CommercialForSalePropertyWithManagerResponseDto{
		Id:               commercialPropertyForSale.Id,
		ManagerId:        commercialPropertyForSale.ManagerId,
		UniqueId:         commercialPropertyForSale.UniqueId,
		Price:            commercialPropertyForSale.Price,
		SizeNumber:       commercialPropertyForSale.SizeNumber,
		SizeDimensions:   commercialPropertyForSale.SizeDimensions,
		Status:           commercialPropertyForSale.Status,
		Type:             commercialPropertyForSale.Type,
		YearBuilt:        commercialPropertyForSale.YearBuilt,
		Stories:          commercialPropertyForSale.Stories,
		HasElectricity:   commercialPropertyForSale.HasElectricity,
		IsNegotiable:     commercialPropertyForSale.IsNegotiable,
		HasWater:         commercialPropertyForSale.HasWater,
		NumberOfRooms:    commercialPropertyForSale.NumberOfRooms,
		OtherDetails:     commercialPropertyForSale.OtherDetails,
		ExteriorFeatures: commercialPropertyForSale.ExteriorFeatures,
		InteriorFeatures: commercialPropertyForSale.InteriorFeatures,
		PostedTime:       generalUtilities.GetTimePassed(commercialPropertyForSale.CreatedAt),
		PropertyLocation: PropertyLocationResponse(commercialPropertyForSale.Location),
		Insights:         PropertyInsightsResponse(commercialPropertyForSale.PropertyInsights),
		Media:            ProcessedPropertyImageAndVideosListToResponse(commercialPropertyForSale.PropertyMedia),
		Manager:          managerUtilities.ManagerResponse(&commercialPropertyForSale.Manager),
	}
}

func CommercialPropertyForSaleResponse(commercialPropertyForSale managerModels.CommercialForSaleProperty) commercialPropertyDtos.CommercialForSalePropertyResponseDto {
	return commercialPropertyDtos.CommercialForSalePropertyResponseDto{
		Id:               commercialPropertyForSale.Id,
		ManagerId:        commercialPropertyForSale.ManagerId,
		UniqueId:         commercialPropertyForSale.UniqueId,
		Price:            commercialPropertyForSale.Price,
		SizeNumber:       commercialPropertyForSale.SizeNumber,
		SizeDimensions:   commercialPropertyForSale.SizeDimensions,
		Status:           commercialPropertyForSale.Status,
		Type:             commercialPropertyForSale.Type,
		YearBuilt:        commercialPropertyForSale.YearBuilt,
		Stories:          commercialPropertyForSale.Stories,
		HasElectricity:   commercialPropertyForSale.HasElectricity,
		IsNegotiable:     commercialPropertyForSale.IsNegotiable,
		HasWater:         commercialPropertyForSale.HasWater,
		NumberOfRooms:    commercialPropertyForSale.NumberOfRooms,
		OtherDetails:     commercialPropertyForSale.OtherDetails,
		ExteriorFeatures: commercialPropertyForSale.ExteriorFeatures,
		InteriorFeatures: commercialPropertyForSale.InteriorFeatures,
		PostedTime:       generalUtilities.GetTimePassed(commercialPropertyForSale.CreatedAt),
		PropertyLocation: PropertyLocationResponse(commercialPropertyForSale.Location),
		Insights:         PropertyInsightsResponse(commercialPropertyForSale.PropertyInsights),
		Media:            ProcessedPropertyImageAndVideosListToResponse(commercialPropertyForSale.PropertyMedia),
	}
}
