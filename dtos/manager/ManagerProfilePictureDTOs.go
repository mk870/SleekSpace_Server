package dtos

type ManagerProfilePictureCreationDTO struct {
	Uri         string `json:"uri"`
	Name        string `json:"name"`
	FullPath    string `json:"fullPath"`
	Size        int    `json:"size"`
	ContentType string `json:"contentType"`
	FileType    string `json:"fileType"`
}

type ManagerProfilePictureResponseAndUpdateDTO struct {
	Id          int    `json:"id"`
	ManagerId   int    `json:"managerId"`
	Uri         string `json:"uri"`
	Name        string `json:"name"`
	FullPath    string `json:"fullPath"`
	Size        int    `json:"size"`
	ContentType string `json:"contentType"`
	FileType    string `json:"fileType"`
}
