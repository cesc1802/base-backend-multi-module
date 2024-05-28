package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string
	LoginID   string
	Password  string
	Salt      string
	Status    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *gorm.DeletedAt
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) IsActive() bool {
	return true
}
