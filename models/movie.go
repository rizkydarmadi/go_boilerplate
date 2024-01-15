package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Name   string `json:"name" gorm:"type:varchar(200);"`
	Title  string `json:"title" gorm:"type:varchar(200);"`
	Price  uint   `json:"price" gorm:"type:uint;"`
	UserID uint
	User   User `gorm:"foreignkey:UserID"`
}
