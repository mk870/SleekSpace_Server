package media

import (
	"net/http"

	imagesOrVideosDtos "SleekSpace/dtos/property/media"
	propertyModels "SleekSpace/models/property"
	imagesOrVideosRepo "SleekSpace/repositories/property/media"
	generalServices "SleekSpace/services/property"
	propertyUtilities "SleekSpace/utilities/funcs/property"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreatePropertyImageOrVideoWithPropertyId(c *gin.Context) {
	var imageOrVideoInfo imagesOrVideosDtos.PropertyImageOrVideoCreationWithPropertyIdDto
	validateModelFields := validator.New()
	c.BindJSON(&imageOrVideoInfo)

	modelFieldsValidationError := validateModelFields.Struct(imageOrVideoInfo)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	newImageOrVideo := propertyModels.PropertyImageOrVideo{
		PropertyId:   imageOrVideoInfo.PropertyId,
		FullPath:     imageOrVideoInfo.FullPath,
		Size:         imageOrVideoInfo.Size,
		FileType:     imageOrVideoInfo.FileType,
		ContentType:  imageOrVideoInfo.ContentType,
		Uri:          imageOrVideoInfo.Uri,
		Name:         imageOrVideoInfo.Name,
		PropertyType: imageOrVideoInfo.PropertyType,
	}

	isImageOrVideoCreated := imagesOrVideosRepo.CreatePropertyImageOrVideo(&newImageOrVideo)
	if !isImageOrVideoCreated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to post your media file"})
		return
	}
	generalServices.GetPropertyTypeById(c, imageOrVideoInfo.PropertyType, imageOrVideoInfo.PropertyId)
}

func GetPropertyImageOrVideoById(c *gin.Context) {
	propertyImageOrVideo := imagesOrVideosRepo.GetPropertyImageOrVideoById(c.Param("id"))
	if propertyImageOrVideo == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this media file does not exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": propertyUtilities.PropertyImageOrVideoResponse(*propertyImageOrVideo)})
}

func UpdatePropertyImageOrVideo(c *gin.Context) {
	var imageOrVideoUpdateDetails imagesOrVideosDtos.PropertyImageOrVideoUpdateAndResponseDto
	validateModelFields := validator.New()
	c.BindJSON(&imageOrVideoUpdateDetails)

	modelFieldsValidationError := validateModelFields.Struct(imageOrVideoUpdateDetails)
	if modelFieldsValidationError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": modelFieldsValidationError.Error()})
		return
	}

	imageOrVideoUpdate := propertyModels.PropertyImageOrVideo{
		Id:           imageOrVideoUpdateDetails.Id,
		PropertyId:   imageOrVideoUpdateDetails.PropertyId,
		Size:         imageOrVideoUpdateDetails.Size,
		FullPath:     imageOrVideoUpdateDetails.FullPath,
		FileType:     imageOrVideoUpdateDetails.FileType,
		ContentType:  imageOrVideoUpdateDetails.ContentType,
		Uri:          imageOrVideoUpdateDetails.Uri,
		Name:         imageOrVideoUpdateDetails.Name,
		PropertyType: imageOrVideoUpdateDetails.PropertyType,
	}

	isImageOrVideoUpdated := imagesOrVideosRepo.UpdatePropertyImageOrVideo(&imageOrVideoUpdate)
	if !isImageOrVideoUpdated {
		c.JSON(http.StatusForbidden, gin.H{"error": "failed to update your media file"})
		return
	}
	generalServices.GetPropertyTypeById(c, imageOrVideoUpdateDetails.PropertyType, imageOrVideoUpdate.PropertyId)
}

func DeletePropertyImageOrVideo(c *gin.Context) {
	imageOrVideo := imagesOrVideosRepo.GetPropertyImageOrVideoById(c.Param("id"))
	if imageOrVideo == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "this media file does not exist"})
		return
	}
	isImageOrVideoDeleted := imagesOrVideosRepo.DeletePropertyImageOrVideo(c.Param("id"))
	if !isImageOrVideoDeleted {
		c.JSON(http.StatusForbidden, gin.H{"error": "this media file does not exist"})
		return
	} else {
		generalServices.GetPropertyTypeById(c, imageOrVideo.PropertyType, imageOrVideo.PropertyId)
	}
}
