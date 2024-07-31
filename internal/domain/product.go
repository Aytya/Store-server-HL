package domain

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null" validate:"required"`
	Description string    `gorm:"not null" validate:"required"`
	Price       float64   `gorm:"not null" validate:"required,gt=0"`
	Category    string    `gorm:"not null" validate:"required"`
	Quantity    int       `gorm:"not null" json:"quantity" validate:"required,gte=0"`
	CreatedAt   time.Time `gorm:"not null;autoCreateTime"`
}

var ProductBaseMessages = map[string]string{
	"required": "is required",
	"gt":       "must be greater than 0",
	"gte":      "must be greater than or equal to 0",
}
