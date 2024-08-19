package imageorvideo

type PropertyImageOrVideoCreationWithNoPropertyIdDto struct {
	Uri          string `json:"uri"`
	Name         string `json:"name"`
	FullPath     string `json:"fullPath"`
	Size         int    `json:"size"`
	ContentType  string `json:"contentType"`
	FileType     string `json:"fileType"`
	PropertyType string `json:"propertyType"`
}

type PropertyImageOrVideoCreationWithPropertyIdDto struct {
	PropertyId   int    `json:"propertyId"`
	Uri          string `json:"uri"`
	Name         string `json:"name"`
	FullPath     string `json:"fullPath"`
	Size         int    `json:"size"`
	ContentType  string `json:"contentType"`
	FileType     string `json:"fileType"`
	PropertyType string `json:"propertyType"`
}

type PropertyImageOrVideoUpdateAndResponseDto struct {
	Id           int    `json:"id"`
	PropertyId   int    `json:"propertyId"`
	Uri          string `json:"uri"`
	Name         string `json:"name"`
	FullPath     string `json:"fullPath"`
	Size         int    `json:"size"`
	ContentType  string `json:"contentType"`
	FileType     string `json:"fileType"`
	PropertyType string `json:"propertyType"`
}
