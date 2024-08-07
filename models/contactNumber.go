package models

type ContactNumber struct {
	MyModel
	Id           int    `json:"id" gorm:"primary_key"`
	UserId       int    `json:"userId"`
	Type         string `json:"type"`
	Number       string `json:"number"`
	CountryCode  string `json:"countryCode"`
	CountryAbbrv string `json:"countryAbbrv"`
}
