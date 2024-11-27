package entities

import (
	"errors"
	"time"
)

var (
	ErrorProductNameEmpty = errors.New("product name cannot be empty")
	ErrorInvalidPrice     = errors.New("product price must be greater than 0")
)

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Price     int    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return ErrorProductNameEmpty
	}
	if p.Price <= 0 {
		return ErrorInvalidPrice
	}

	return nil
}

func (Product) TableName() string {
	return "products"
}
