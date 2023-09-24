package model

import "time"

type Item struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	ItemCode    string    `json:"item_code" gorm:"column:item_code;not null"`
	Description string    `json:"description" gorm:"type:text"`
	Quantity    int       `json:"quantity" gorm:"not null"`
	OrderID     uint      `json:"order_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
