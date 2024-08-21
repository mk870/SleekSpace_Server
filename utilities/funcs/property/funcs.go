package property

import (
	commercialPropertyDtos "SleekSpace/dtos/property/commercial"
	propertyInsightsDtos "SleekSpace/dtos/property/insights"
	landDtos "SleekSpace/dtos/property/land"
	propertyLocationDtos "SleekSpace/dtos/property/location"
	propertyMediaDtos "SleekSpace/dtos/property/media"
	residentialPropertyDtos "SleekSpace/dtos/property/residential"
	propertyStandDtos "SleekSpace/dtos/property/stand"
	propertyModels "SleekSpace/models/property"
	"math/rand"
	"time"
)

func GeneratePropertyUniqueId() int {
	rand.Seed(time.Now().UnixNano())
	min := 1000000000000000
	max := 9999999999999999
	randomInt := rand.Intn(max-min) + min
	return randomInt
}

func ConvertPropertyImagesOrVideosWithNoPropertyIdToModel(propertyMediaList []propertyMediaDtos.PropertyImageOrVideoCreationWithNoPropertyIdDto, propertyType string) []propertyModels.PropertyImageOrVideo {
	mediaList := []propertyModels.PropertyImageOrVideo{}
	if len(propertyMediaList) > 0 {
		for i := 0; i < len(propertyMediaList); i++ {
			media := propertyModels.PropertyImageOrVideo{
				Uri:          propertyMediaList[i].Uri,
				FileType:     propertyMediaList[i].FileType,
				ContentType:  propertyMediaList[i].ContentType,
				Size:         propertyMediaList[i].Size,
				Name:         propertyMediaList[i].Name,
				FullPath:     propertyMediaList[i].FullPath,
				PropertyType: propertyType,
			}
			mediaList = append(mediaList, media)
		}
	}
	return mediaList
}

func ProcessedPropertyImageAndVideosListToResponse(mediaList []propertyModels.PropertyImageOrVideo) []propertyMediaDtos.PropertyImageOrVideoUpdateAndResponseDto {
	dtoList := []propertyMediaDtos.PropertyImageOrVideoUpdateAndResponseDto{}
	if len(mediaList) > 0 {
		for i := 0; i < len(mediaList); i++ {
			dto := propertyMediaDtos.PropertyImageOrVideoUpdateAndResponseDto{
				Id:           mediaList[i].Id,
				PropertyId:   mediaList[i].PropertyId,
				Uri:          mediaList[i].Uri,
				FileType:     mediaList[i].FileType,
				ContentType:  mediaList[i].ContentType,
				Size:         mediaList[i].Size,
				Name:         mediaList[i].Name,
				FullPath:     mediaList[i].FullPath,
				PropertyType: mediaList[i].PropertyType,
			}
			dtoList = append(dtoList, dto)
		}
	}
	return dtoList
}

func PropertyImageOrVideoResponse(media propertyModels.PropertyImageOrVideo) propertyMediaDtos.PropertyImageOrVideoUpdateAndResponseDto {
	return propertyMediaDtos.PropertyImageOrVideoUpdateAndResponseDto{
		Id:           media.Id,
		PropertyId:   media.PropertyId,
		Uri:          media.Uri,
		FileType:     media.FileType,
		ContentType:  media.ContentType,
		Size:         media.Size,
		Name:         media.Name,
		FullPath:     media.FullPath,
		PropertyType: media.PropertyType,
	}
}

func PropertyInsightsResponse(insights propertyModels.PropertyInsights) propertyInsightsDtos.PropertyInsightsUpdateAndResponseDto {
	return propertyInsightsDtos.PropertyInsightsUpdateAndResponseDto{
		Views:             insights.Views,
		Id:                insights.Id,
		PropertyId:        insights.PropertyId,
		ContactInfoViews:  insights.ContactInfoViews,
		AddedToFavourites: insights.AddedToFavourites,
		Shared:            insights.Shared,
		PropertyType:      insights.PropertyType,
	}
}

func PropertyLocationResponse(location propertyModels.PropertyLocation) propertyLocationDtos.PropertyLocationUpdateAndResponseDto {
	return propertyLocationDtos.PropertyLocationUpdateAndResponseDto{
		Id:           location.Id,
		PropertyId:   location.PropertyId,
		Boundingbox:  location.Boundingbox,
		Lat:          location.Lat,
		Lon:          location.Lon,
		DisplayName:  location.DisplayName,
		City:         location.City,
		County:       location.County,
		Country:      location.Country,
		CountryCode:  location.CountryCode,
		Province:     location.Province,
		Surburb:      location.Surburb,
		PropertyType: location.PropertyType,
	}
}

func PropertyStandResponse(standModel propertyModels.PropertyStand) propertyStandDtos.StandResponseDTO {
	return propertyStandDtos.StandResponseDTO{
		Id:                 standModel.Id,
		ManagerId:          standModel.ManagerId,
		UniqueId:           standModel.UniqueId,
		Price:              standModel.Price,
		SizeNumber:         standModel.SizeNumber,
		SizeDimensions:     standModel.SizeDimensions,
		Status:             standModel.Status,
		IsServiced:         standModel.IsServiced,
		IsNegotiable:       standModel.IsNegotiable,
		AreaHasElectricity: standModel.AreaHasElectricity,
		Level:              standModel.Level,
		Type:               standModel.Type,
		PropertyLocation:   PropertyLocationResponse(standModel.Location),
		Insights:           PropertyInsightsResponse(standModel.PropertyInsights),
		Media:              ProcessedPropertyImageAndVideosListToResponse(standModel.PropertyMedia),
	}
}

func ResidentialRentalPropertyResponse(residentialRentalModel propertyModels.ResidentialRentalProperty) residentialPropertyDtos.ResidentialPropertyForRentResponseDto {
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
		PropertyLocation: PropertyLocationResponse(residentialRentalModel.Location),
		Insights:         PropertyInsightsResponse(residentialRentalModel.PropertyInsights),
		Media:            ProcessedPropertyImageAndVideosListToResponse(residentialRentalModel.PropertyMedia),
	}
}

func ResidentialForSalePropertyResponse(residentialForSaleModel propertyModels.ResidentialPropertyForSale) residentialPropertyDtos.ResidentialPropertyForSaleResponseDto {
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
		PropertyLocation: PropertyLocationResponse(residentialForSaleModel.Location),
		Insights:         PropertyInsightsResponse(residentialForSaleModel.PropertyInsights),
		Media:            ProcessedPropertyImageAndVideosListToResponse(residentialForSaleModel.PropertyMedia),
	}
}

func CommercialPropertyForSaleResponse(commercialPropertyForSale propertyModels.CommercialForSaleProperty) commercialPropertyDtos.CommercialForSalePropertyResponseDto {
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
		PropertyLocation: PropertyLocationResponse(commercialPropertyForSale.Location),
		Insights:         PropertyInsightsResponse(commercialPropertyForSale.PropertyInsights),
		Media:            ProcessedPropertyImageAndVideosListToResponse(commercialPropertyForSale.PropertyMedia),
	}
}

func CommercialPropertyForRentResponse(commercialPropertyForRent propertyModels.CommercialRentalProperty) commercialPropertyDtos.CommercialForRentPropertyResponseDto {
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
		PropertyLocation: PropertyLocationResponse(commercialPropertyForRent.Location),
		Insights:         PropertyInsightsResponse(commercialPropertyForRent.PropertyInsights),
		Media:            ProcessedPropertyImageAndVideosListToResponse(commercialPropertyForRent.PropertyMedia),
	}
}

func LandPropertyResponse(landPropertyModel propertyModels.LandForSaleProperty) landDtos.LandForSalePropertyResponseDto {
	return landDtos.LandForSalePropertyResponseDto{
		Id:                 landPropertyModel.Id,
		ManagerId:          landPropertyModel.ManagerId,
		UniqueId:           landPropertyModel.UniqueId,
		Price:              landPropertyModel.Price,
		SizeNumber:         landPropertyModel.SizeNumber,
		SizeDimensions:     landPropertyModel.SizeDimensions,
		Status:             landPropertyModel.Status,
		Type:               landPropertyModel.Type,
		AreaHasElectricity: landPropertyModel.AreaHasElectricity,
		HasWater:           landPropertyModel.HasWater,
		IsNegotiable:       landPropertyModel.IsNegotiable,
		PropertyLocation:   PropertyLocationResponse(landPropertyModel.Location),
		Insights:           PropertyInsightsResponse(landPropertyModel.PropertyInsights),
		Media:              ProcessedPropertyImageAndVideosListToResponse(landPropertyModel.PropertyMedia),
	}
}
