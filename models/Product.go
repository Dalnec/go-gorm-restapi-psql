package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Description string `gorm:"not null" json:"description"`
	Active      bool   `gorm:"default:true" json:"active"`
	Product     []Product `json:"products"`
}

type Brand struct {
	gorm.Model
	Description string `gorm:"not null" json:"description"`
	Active      bool   `gorm:"default:true" json:"active"`
	Product     []Product `json:"products"`
}

type Product struct {
	gorm.Model
	Code       	string `gorm:"not null;unique_index" json:"code"`
	Description string `gorm:"not null" json:"description"`
	Price 		float32 `sql:"type:decimal(10,2);" json:"price"`
	MinPrice 	float32 `sql:"type:decimal(10,2);" json:"minprice"`
	Active      bool   `gorm:"default:true" json:"active"`

	Category   Category `json:"category"`
	Brand      Brand    `json:"brand"`
	User      User    `json:"user"`

	CategoryID  uint   `json:"category_id"`
	BrandID     uint   `json:"brand_id"`
	UserID      uint   `json:"user_id"`
}
// image