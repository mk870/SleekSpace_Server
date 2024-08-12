package models

type ManagerContactNumber struct {
	MyModel
	Id           int    `json:"id" gorm:"primary_key"`
	ManagerId    int    `json:"managerId" gorm:"column:managerId"`
	Type         string `json:"type"`
	Number       string `json:"number"`
	CountryCode  string `json:"countryCode"`
	CountryAbbrv string `json:"countryAbbrv"`
}
