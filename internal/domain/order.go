package domain

import "time"

type Order struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `json:"user_id" validate:"required"`
	ProductIDs []uint    `gorm:"-" json:"product_ids" validate:"required"`
	TotalPrice float64   `json:"total_price" validate:"required"`
	OrderDate  time.Time `gorm:"not null;autoCreateTime"`
	Status     string    `validate:"required,oneof=new processing completed"`
}

var OrderBaseMessages = map[string]string{
	"required": "is required",
	"oneof":    "must be either 'new', 'processing' or 'completed'",
}
