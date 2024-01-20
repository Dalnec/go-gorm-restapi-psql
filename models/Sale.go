package models

import (
	"gorm.io/gorm"
)


type PaymentMethod struct {
	gorm.Model
	Description string `gorm:"not null" json:"description"`
	Active      bool   `gorm:"default:true" json:"active"`
	Sale     []Sale `json:"sales"`
}

type Sale struct {
	gorm.Model
	IssueDate 	string `gorm:"not null" json:"issuedate"`
	Number      string `gorm:"not null;unique_index" json:"number"`
	Customer 	string `gorm:"not null" json:"customer"`
	Total		float32 `sql:"type:decimal(10,2);" json:"total"`
	Active      bool   `gorm:"default:true" json:"active"`

	PaymentMethod   PaymentMethod `json:"paymentmethod"`
	User       		User     `json:"user"`

	PaymentMethodID  uint   `json:"paymentmethod_id"`
	UserID      	 uint   `json:"user_id"`

	SaleDetail     []SaleDetail `json:"details"`
}

type SaleDetail struct {
	gorm.Model
	Product		Product `json:"product"`
	Prices		Prices `json:"´prices"`
	Sale		Sale `json:"sale"`
	Quantity	float32 `sql:"type:decimal(10,2);" json:"quantity"`
	Total		float32 `sql:"type:decimal(10,2);" json:"total"`

	PricesID	uint   `json:"prices_id"`
	ProductID	uint   `json:"product_id"`
	SaleID		uint   `json:"sale_id"`
}