package dtos

type UserResponseDTO struct {
	Email         string             `json:"email"`
	Id            int                `json:"id"`
	ContactNumber []ContactNumberDTO `json:"contactNumber"`
	Avatar        string             `json:"avatar"`
	Location      LocationDTO        `json:"location"`
	GivenName     string             `json:"givenName"`
	FamilyName    string             `json:"familyName"`
	AccessToken   string             `json:"accessToken"`
}

type ContactNumberDTO struct {
	Id           int    `json:"id"`
	UserId       int    `json:"userId"`
	Type         string `json:"type"`
	Number       string `json:"number"`
	CountryCode  string `json:"countryCode"`
	CountryAbbrv string `json:"countryAbbrv"`
}
type LocationDTO struct {
	Id          int    `json:"id"`
	UserId      int    `json:"userId"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
	Surburb     string `json:"countryAbbrv"`
	City        string `json:"city"`
	County      string `json:"county"`
	Province    string `json:"province"`
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
}
