package dtos

type UserResponseDTO struct {
	Email          string `json:"email"`
	Id             int    `json:"id"`
	WhatsAppNumber string `json:"whatsAppNumber"`
	ContactNumber  string `json:"contactNumber"`
	Avatar         string `json:"avatar"`
	Location       string `json:"location"`
	GivenName      string `json:"givenName"`
	FamilyName     string `json:"familyName"`
	AccessToken    string `json:"accessToken"`
}
