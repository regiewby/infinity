package model

import "time"

type UserCreateRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	PIN         string `json:"pin"`
}

type UserResponse struct {
	ID          uint64    `json:"id,omitempty"`
	SafeID      string    `json:"safeID,omitempty"`
	Email       string    `json:"email,omitempty"`
	PhoneNumber string    `json:"phoneNumber,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
