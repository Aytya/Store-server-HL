package domain

import "time"

type Payment struct {
	ID            uint      `gorm:"primaryKey"`
	UserID        uint      `json:"user_id" gorm:"not null"`
	OrderID       uint      `json:"order_id" gorm:"not null"`
	Amount        float64   `gorm:"not null"`
	PaymentDate   time.Time `gorm:"autoCreateTime"`
	PaymentStatus string    `json:"payment_status" gorm:"-"`
}
