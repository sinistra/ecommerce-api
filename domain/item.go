package domain

import (
	"time"
)

type Item struct {
	Id                int       `json:"id" db:"title"`
	Title             string    `json:"title" db:"title"`
	Description       string    `json:"description" db:"description"`
	Seller            int64     `json:"seller" db:"seller"`
	Picture           string    `json:"pictures" db:"pictures"`
	Price             float64   `json:"price" db:"price"`
	AvailableQuantity int       `json:"available_quantity" db:"available_quantity"`
	SoldQuantity      int       `json:"sold_quantity" db:"sold_quantity"`
	Status            string    `json:"status" db:"status"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}
