package models

import (
	"time"
)

type Product struct {
	ID        uint       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
	Name      string     `json:"name" validate:"required"`
	Price     float64    `json:"price" validate:"required"`
	Stock     int        `json:"stock" validate:"required"`
}

type ProductUpdate struct {
	Name  *string  `json:"name"`
	Price *float64 `json:"price"`
	Stock *int     `json:"stock"`
}
