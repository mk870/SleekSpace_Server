package models

type Location struct {
	MyModel
	Id          int    `json:"id" gorm:"primary_key"`
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
