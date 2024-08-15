package models

import baseModel "SleekSpace/models"

type PropertyMetrics struct {
	baseModel.MyModel
	Id         int `json:"id" gorm:"primary_key"`
	PropertyId int `json:"propertyId" gorm:"column:propertyId"`
	Views      int `json:"views"`
	Likes      int `json:"likes"`
}
