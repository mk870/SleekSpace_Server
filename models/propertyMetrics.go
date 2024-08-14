package models

type PropertyMetrics struct {
	MyModel
	Id         int `json:"id" gorm:"primary_key"`
	PropertyId int `json:"propertyId" gorm:"column:propertyId"`
	Views      int `json:"views"`
	Likes      int `json:"likes"`
}
