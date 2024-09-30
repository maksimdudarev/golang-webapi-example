package models

import "gorm.io/gorm"

type Discount struct {
	gorm.Model
	ProductName string
	Description string
	Amount      int
}
