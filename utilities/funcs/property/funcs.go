package property

import (
	propertyMediaDtos "SleekSpace/dtos/property/imageorvideo"
	propertyInsightsDtos "SleekSpace/dtos/property/insights"
	propertyLocationDtos "SleekSpace/dtos/property/location"
	propertyStandDtos "SleekSpace/dtos/property/stand"
	propertyModels "SleekSpace/models/property"
)

func EmptyPropertyInsights() propertyModels.PropertyInsights {
	return propertyModels.PropertyInsights{
		Views:             0,
		Shared:            0,
		AddedToFavourites: 0,
		ContactInfoViews:  0,
		PropertyType:      "",
	}
}

func ConvertPropertyImagesOrVideosWithNoPropertyIdToModel(propertyMediaList []propertyMediaDtos.PropertyImageOrVideoCreationWithNoPropertyIdDto) []propertyModels.PropertyImageOrVideo {
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
				PropertyType: propertyMediaList[i].PropertyType,
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
