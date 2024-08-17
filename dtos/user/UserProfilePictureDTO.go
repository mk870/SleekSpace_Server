package dtos

type UserProfilePictureResponseAndUpdateDTO struct {
	Id          int    `json:"id"`
	UserId      int    `json:"userId"`
	Uri         string `json:"uri"`
	Name        string `json:"name"`
	FullPath    string `json:"fullPath"`
	Size        int    `json:"size"`
	ContentType string `json:"contentType"`
	FileType    string `json:"fileType"`
}
type UserProfilePictureCreationDTO struct {
	UserId      int    `json:"userId"`
	Uri         string `json:"uri"`
	Name        string `json:"name"`
	FullPath    string `json:"fullPath"`
	Size        int    `json:"size"`
	ContentType string `json:"contentType"`
	FileType    string `json:"fileType"`
}
