package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
	UserName  string `gorm:"not null" json:"user_name"`
	Product     []Product `json:"products"`
}
