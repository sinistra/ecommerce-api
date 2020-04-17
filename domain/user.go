package domain

import (
    "time"
)

// User describes a model
type User struct {
    Id          int       `json:"id"`
    FirstName   string    `json:"first_name"`
    LastName    string    `json:"last_name"`
    Email       string    `json:"email" binding:"required"`
    DateCreated time.Time `json:"date_created"`
    Status      string    `json:"status"`
    Password    string    `json:"password" binding:"required"`
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
