package model

import (
	"strings"
	"time"

	"gorm.io/gorm"
	"github.com/ldhoax/SOS-requester/utils"
	"github.com/google/uuid"
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
	r.Username = strings.TrimSpace(r.ID)
	return
}

func (r *Requester) BeforeSave(tx *gorm.DB) (err error) {
	r.Username = strings.TrimSpace(r.Username)

	hashedPassword, err := utils.HashPassword(r.ID + "_" + r.CreatedAt.Format("20060102150405"))
	if err != nil {
		panic(err)
	}
	r.Password = string(hashedPassword)
	return
}
