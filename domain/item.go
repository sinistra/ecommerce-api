package domain

import (
    "time"
)

type Item struct {
    Id                int       `json:"id"`
    Title             string    `json:"title"`
    Description       string    `json:"description"`
    Seller            int64     `json:"seller"`
    Picture           string    `json:"pictures"`
    Price             float64   `json:"price"`
    AvailableQuantity int       `json:"available_quantity"`
    SoldQuantity      int       `json:"sold_quantity"`
    Status            string    `json:"status"`
    CreatedAt         time.Time `json:" created_at"`
    UpdatedAt         time.Time `json:" updated_at"`
}
