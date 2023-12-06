package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id" required:"true"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	IsActive    bool      `json:"is_active"`
}
