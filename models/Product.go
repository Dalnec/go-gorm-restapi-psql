package models

import "gorm.io/gorm"

type Measure struct {
	gorm.Model
	Description string `gorm:"not null" json:"description"`
	Active      bool   `gorm:"default:true" json:"active"`
	Prices     []Prices `json:"prices"`
}

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
	Active      bool   `gorm:"default:true" json:"active"`

	Category   Category `json:"category"`
	Brand      Brand    `json:"brand"`
	User       User     `json:"user"`
	Product    *Product `gorm:"default:null" json:"product"`

	CategoryID  uint   `json:"category_id"`
	BrandID     uint   `json:"brand_id"`
	UserID      uint   `json:"user_id"`
	ProductID   *uint   `gorm:"default:null" json:"product_id"`

	Prices     []Prices `json:"prices"`
	Products     []Product `gorm:"foreignKey:ID" json:"products"`
}

type Prices struct {
	gorm.Model
	Equivalent	uint8 `gorm:"not null" json:"equivalent"`
	Price		float32 `sql:"type:decimal(10,2);" json:"price"`
	MinPrice 	float32 `sql:"type:decimal(10,2);" json:"minprice"`
	Active      bool   `gorm:"default:true" json:"active"`

	Product    Product `json:"product"`
	Measure    Measure `json:"measure"`

	ProductID   uint   `json:"product_id"`
	MeasureID	uint   `json:"measure_id"`
}