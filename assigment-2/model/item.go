package model

import (
	"time"
)

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	ItemCode    string `json:"itemcode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint   `json:"orderId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
