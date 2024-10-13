package model

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
	"time"
)

type Requester struct {
	ID        string    `gorm:"type:uuid;primary_key;"`
	Email     string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

func (r *Requester) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.New().String()
	r.Username = r.ID
	r.Password = r.ID + "_" + r.CreatedAt.Format("20060102150405")
	return
}