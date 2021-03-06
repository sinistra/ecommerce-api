package domain

import (
	"time"
)

// User describes the user model
type User struct {
	Id        int       `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" binding:"required" db:"email"`
	CreatedAt time.Time `json:"created_at, omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at, omitempty" db:"updated_at"`
	UUID      *string   `json:"uuid, omitempty" db:"uuid"`
	Status    string    `json:"status, omitempty" db:"status"`
	Password  string    `json:"password, omitempty" db:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
