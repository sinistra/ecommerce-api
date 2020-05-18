package domain

import (
	"time"
)

type Item struct {
	Id                int       `json:"id" db:"id"`
	Code              string    `json:"code" db:"code"`
	Description       string    `json:"description" db:"description"`
	Seller            int64     `json:"seller" db:"seller"`
	Picture           string    `json:"picture" db:"picture"`
	Price             float64   `json:"price" db:"price"`
	AvailableQuantity int       `json:"qty_avail" db:"qty_avail"`
	SoldQuantity      int       `json:"qty_sold" db:"qty_sold"`
	Status            string    `json:"status" db:"status"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}
