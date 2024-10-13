package model

import (
	"html"
	"strings"
	"time"

	"github.com/ldhoax/SOS-requester/utils"
	"gorm.io/gorm"
)

type Status int

const (
	StatusActive Status = iota + 1
	StatusDeactive
)

func (s Status) IsValid() bool {
	switch s {
	case StatusActive:
		return true
	case StatusDeactive:
		return true
	}
	return false
}

type User struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Password  string    `gorm:"size:255;not null" json:"password"`
	Email     string    `gorm:"size:255;not null" json:"email"`
	Status    Status    `gorm:"embedded" json:"status"`
	ID        uint      `json:"id"`
}

func (u *User) BeforeSave(db *gorm.DB) error {

	//turn password into hash
	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}
