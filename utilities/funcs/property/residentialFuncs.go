package property

import (
	residentialPropertyDtos "SleekSpace/dtos/property/residential"
	managerModels "SleekSpace/models/manager"
	generalUtilities "SleekSpace/utilities/funcs/general"
	managerUtilities "SleekSpace/utilities/funcs/manager"
)

func ResidentialRentalPropertyWithManagerResponse(residentialRentalModel managerModels.ResidentialRentalProperty) residentialPropertyDtos.ResidentialPropertyForRentWithManagerResponseDto {
	return residentialPropertyDtos.ResidentialPropertyForRentWithManagerResponseDto{
		Id:               residentialRentalModel.Id,
		ManagerId:        residentialRentalModel.ManagerId,
		UniqueId:         residentialRentalModel.UniqueId,
		RentAmount:       residentialRentalModel.RentAmount,
		SizeNumber:       residentialRentalModel.SizeNumber,
		SizeDimensions:   residentialRentalModel.SizeDimensions,
		Status:           residentialRentalModel.Status,
		Type:             residentialRentalModel.Type,
		YearBuilt:        residentialRentalModel.YearBuilt,
		Bedrooms:         residentialRentalModel.Bedrooms,
		Bathrooms:        residentialRentalModel.Bathrooms,
		Stories:          residentialRentalModel.Stories,
		HasElectricity:   residentialRentalModel.HasElectricity,
		HasWater:         residentialRentalModel.HasWater,
		IsFullHouse:      residentialRentalModel.IsFullHouse,
		NumberOfRooms:    residentialRentalModel.NumberOfRooms,
		NumberOfGarages:  residentialRentalModel.NumberOfGarages,
		HasSwimmingPool:  residentialRentalModel.HasSwimmingPool,
		OtherDetails:     residentialRentalModel.OtherDetails,
		ExteriorFeatures: residentialRentalModel.ExteriorFeatures,
		InteriorFeatures: residentialRentalModel.InteriorFeatures,
		PostedTime:       generalUtilities.GetTimePassed(residentialRentalModel.CreatedAt),
		PropertyLocation: PropertyLocationResponse(residentialRentalModel.Location),
		Insights:         PropertyInsightsResponse(residentialRentalModel.PropertyInsights),
		Media:            ProcessedPropertyImageAndVideosListToResponse(residentialRentalModel.PropertyMedia),
		Manager:          managerUtilities.ManagerResponse(&residentialRentalModel.Manager),
	}
}

func ResidentialForSalePropertyWithManagerResponse(residentialForSaleModel managerModels.ResidentialPropertyForSale) residentialPropertyDtos.ResidentialPropertyForSaleWithManagerResponseDto {
	return residentialPropertyDtos.ResidentialPropertyForSaleWithManagerResponseDto{
		Id:               residentialForSaleModel.Id,
		ManagerId:        residentialForSaleModel.ManagerId,
		UniqueId:         residentialForSaleModel.UniqueId,
		Price:            residentialForSaleModel.Price,
		SizeNumber:       residentialForSaleModel.SizeNumber,
		SizeDimensions:   residentialForSaleModel.SizeDimensions,
		Status:           residentialForSaleModel.Status,
		Type:             residentialForSaleModel.Type,
		YearBuilt:        residentialForSaleModel.YearBuilt,
		Bedrooms:         residentialForSaleModel.Bedrooms,
		Bathrooms:        residentialForSaleModel.Bathrooms,
		Stories:          residentialForSaleModel.Stories,
		HasElectricity:   residentialForSaleModel.HasElectricity,
		HasWater:         residentialForSaleModel.HasWater,
		NumberOfRooms:    residentialForSaleModel.NumberOfRooms,
		NumberOfGarages:  residentialForSaleModel.NumberOfGarages,
		HasSwimmingPool:  residentialForSaleModel.HasSwimmingPool,
		IsNegotiable:     residentialForSaleModel.IsNegotiable,
		OtherDetails:     residentialForSaleModel.OtherDetails,
		ExteriorFeatures: residentialForSaleModel.ExteriorFeatures,
		InteriorFeatures: residentialForSaleModel.InteriorFeatures,
		PostedTime:       generalUtilities.GetTimePassed(residentialForSaleModel.CreatedAt),
		PropertyLocation: PropertyLocationResponse(residentialForSaleModel.Location),
		Insights:         PropertyInsightsResponse(residentialForSaleModel.PropertyInsights),
		Media:            ProcessedPropertyImageAndVideosListToResponse(residentialForSaleModel.PropertyMedia),
		Manager:          managerUtilities.ManagerResponse(&residentialForSaleModel.Manager),
	}
}

func ResidentialForSalePropertyResponse(residentialForSaleModel managerModels.ResidentialPropertyForSale) residentialPropertyDtos.ResidentialPropertyForSaleResponseDto {
	return residentialPropertyDtos.ResidentialPropertyForSaleResponseDto{
		Id:               residentialForSaleModel.Id,
		ManagerId:        residentialForSaleModel.ManagerId,
		UniqueId:         residentialForSaleModel.UniqueId,
		Price:            residentialForSaleModel.Price,
		SizeNumber:       residentialForSaleModel.SizeNumber,
		SizeDimensions:   residentialForSaleModel.SizeDimensions,
		Status:           residentialForSaleModel.Status,
		Type:             residentialForSaleModel.Type,
		YearBuilt:        residentialForSaleModel.YearBuilt,
		Bedrooms:         residentialForSaleModel.Bedrooms,
		Bathrooms:        residentialForSaleModel.Bathrooms,
		Stories:          residentialForSaleModel.Stories,
		HasElectricity:   residentialForSaleModel.HasElectricity,
		HasWater:         residentialForSaleModel.HasWater,
		NumberOfRooms:    residentialForSaleModel.NumberOfRooms,
		NumberOfGarages:  residentialForSaleModel.NumberOfGarages,
		HasSwimmingPool:  residentialForSaleModel.HasSwimmingPool,
		IsNegotiable:     residentialForSaleModel.IsNegotiable,
		OtherDetails:     residentialForSaleModel.OtherDetails,
		ExteriorFeatures: residentialForSaleModel.ExteriorFeatures,
		InteriorFeatures: residentialForSaleModel.InteriorFeatures,
		PostedTime:       generalUtilities.GetTimePassed(residentialForSaleModel.CreatedAt),
		PropertyLocation: PropertyLocationResponse(residentialForSaleModel.Location),
		Insights:         PropertyInsightsResponse(residentialForSaleModel.PropertyInsights),
		Media:            ProcessedPropertyImageAndVideosListToResponse(residentialForSaleModel.PropertyMedia),
	}
}

func ResidentialRentalPropertyResponse(residentialRentalModel managerModels.ResidentialRentalProperty) residentialPropertyDtos.ResidentialPropertyForRentResponseDto {
	return residentialPropertyDtos.ResidentialPropertyForRentResponseDto{
		Id:               residentialRentalModel.Id,
		ManagerId:        residentialRentalModel.ManagerId,
		UniqueId:         residentialRentalModel.UniqueId,
		RentAmount:       residentialRentalModel.RentAmount,
		SizeNumber:       residentialRentalModel.SizeNumber,
		SizeDimensions:   residentialRentalModel.SizeDimensions,
		Status:           residentialRentalModel.Status,
		Type:             residentialRentalModel.Type,
		YearBuilt:        residentialRentalModel.YearBuilt,
		Bedrooms:         residentialRentalModel.Bedrooms,
		Bathrooms:        residentialRentalModel.Bathrooms,
		Stories:          residentialRentalModel.Stories,
		HasElectricity:   residentialRentalModel.HasElectricity,
		HasWater:         residentialRentalModel.HasWater,
		IsFullHouse:      residentialRentalModel.IsFullHouse,
		NumberOfRooms:    residentialRentalModel.NumberOfRooms,
		NumberOfGarages:  residentialRentalModel.NumberOfGarages,
		HasSwimmingPool:  residentialRentalModel.HasSwimmingPool,
		OtherDetails:     residentialRentalModel.OtherDetails,
		ExteriorFeatures: residentialRentalModel.ExteriorFeatures,
		InteriorFeatures: residentialRentalModel.InteriorFeatures,
		PostedTime:       generalUtilities.GetTimePassed(residentialRentalModel.CreatedAt),
		PropertyLocation: PropertyLocationResponse(residentialRentalModel.Location),
		Insights:         PropertyInsightsResponse(residentialRentalModel.PropertyInsights),
		Media:            ProcessedPropertyImageAndVideosListToResponse(residentialRentalModel.PropertyMedia),
	}
}
