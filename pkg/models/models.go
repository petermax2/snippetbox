package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var ErrNoRecord = errors.New("models: no matching record found")
var ErrInvalidCredentials = errors.New("models: invalid credentials")
var ErrDuplicateEmail = errors.New("models: duplicate e-mail address")
var ErrAccountDisabled = errors.New("models: account is not active")

type SnippetProvider interface {
	Migrate() error
	Insert(snippet *Snippet) error
	Get(id int) (*Snippet, error)
	Latest() ([]*Snippet, error)
}

type UserProvider interface {
	Migrate() error
	Insert(user *User) error
	Authenticate(email, password string) (uint, error)
	Get(id uint) (*User, error)
}

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
