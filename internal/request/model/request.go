package model

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
	"time"
)

type Request struct {
	ID        string    `gorm:"type:uuid;primary_key;"`
	RequesterID    string    `json:"requester_id"`
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

func (r *Request) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New().String()
	return
}