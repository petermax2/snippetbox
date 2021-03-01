package models

import (
	"time"

	"gorm.io/gorm"
)

type Snippet struct {
	gorm.Model
	CreatedAt time.Time `gorm:"index"`
	Title     string
	Content   string
	Expires   time.Time `gorm:"index"`
}
