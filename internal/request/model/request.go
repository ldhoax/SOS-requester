package model

import (
	"time"
)

type Request struct {
	ID             uint      `json:"id"`
	PhoneNumber    string    `json:"phone_number"`
	Email          string    `json:"email"`
	Location       string    `json:"location"`
	EmergencyLevel int       `json:"emergency_level"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	ShortDescription string    `json:"short_description"`
	Description    string    `json:"description"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`
}