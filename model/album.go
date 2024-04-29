package model

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Title  string  `gorm:"uniqueIndex" json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
