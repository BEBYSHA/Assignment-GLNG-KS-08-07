package model

import "time"

type Order struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	CustomerName string    `json:"customer_name" gorm:"column:customer_name;type:varchar(155);not null"`
	Items        []Item    `json:"items"`
	OrderedAt    string    `json:"ordered_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
