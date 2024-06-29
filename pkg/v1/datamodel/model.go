package datamodel

import (
	"time"

	"gorm.io/gorm"
)

type Family struct {
	ID              string         `json:"id" gorm:"primary_key"`
	Name            string         `json:"name" gorm:"not null"`
	Platform        string         `json:"platform" gorm:"not null"`
	DueDate         int            `json:"due_date" gorm:"not null"`
	PromptPayNumber string         `json:"prompt_pay_number" gorm:"not null"`
	PricesString    string         `json:"-" gorm:"column:prices"`
	Prices          *[]Prices      `json:"prices" gorm:"-"`
	CreatedAt       time.Time      `json:"created_at"`
	CreatedBy       string         `json:"created_by"`
	UpdatedAt       time.Time      `json:"updated_at"`
	UpdatedBy       string         `json:"updated_by"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`
	DeletedBy       string         `json:"-"`
}

type Prices struct {
	Price float64 `json:"price"`
	Month int     `json:"month"`
}
