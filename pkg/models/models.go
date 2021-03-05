package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Snippet struct {
	gorm.Model
	CreatedAt time.Time `gorm:"index"`
	Title     string
	Content   string
	Expires   time.Time `gorm:"index"`
}
