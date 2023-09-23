package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `gorm:"not null" json:"user_name"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	Role      string `gorm:"not null" json:"role"`
	Product     []Product `json:"products"`
}
