package model

import "time"

type Order struct {
	ID uint `gorm:"primaryKey"`
	CustomerName string `json:customerName`
	Items []Item `json:"items"`
	OrderedAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}