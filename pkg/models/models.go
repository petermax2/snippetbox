package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var ErrNoRecord = errors.New("models: no matching record found")
var ErrInvalidCredentials = errors.New("models: invalid credentials")
var ErrDuplicateEmail = errors.New("models: duplicate e-mail address")

type Snippet struct {
	gorm.Model
	CreatedAt time.Time `gorm:"index"`
	Title     string
	Content   string
	Expires   time.Time `gorm:"index"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique;index"`
	Password string
	Active   bool
}
