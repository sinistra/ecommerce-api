package domain

import (
	"time"
)

type Item struct {
	Id                int       `json:"id" db:"id" form:"id"`
	Code              string    `json:"code" db:"code" form:"code"`
	Title             string    `json:"title" db:"title" form:"title"`
	Description       string    `json:"description" db:"description" form:"description"`
	Seller            int64     `json:"seller" db:"seller" form:"seller"`
	Image             string    `json:"image" db:"image" form:"image"`
	Price             float64   `json:"price" db:"price" form:"price"`
	AvailableQuantity int       `json:"qty_avail" db:"qty_avail" form:"qty_avail"`
	SoldQuantity      int       `json:"qty_sold" db:"qty_sold" form:"qty_sold"`
	Status            string    `json:"status" db:"status" form:"status"`
	Featured          bool      `json:"featured" db:"featured" form:"featured"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
	Deleted           bool      `json:"deleted" db:"deleted" form:"deleted"`
}
