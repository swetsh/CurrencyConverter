package models

import (
	"gorm.io/gorm"
)

type Currency struct {
	gorm.Model
	Currency     string `gorm:"unique"`
	ExchangeRate float64
}
