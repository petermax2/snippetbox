package mock

import (
	"time"

	"gorm.io/gorm"
	"nirpet.at/snippetbox/pkg/models"
)

var mockUser = &models.User{
	Model: gorm.Model{
		ID:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	Name:   "Alice",
	Email:  "alice@example.com",
	Active: true,
}

type UserModel struct{}

func (m *UserModel) Migrate() error {
	return nil
}

func (m *UserModel) Insert(user *models.User) error {
	switch user.Email {
	case "duplex@example.com":
		return models.ErrDuplicateEmail
	default:
		user.ID = 2
		user.UpdatedAt = time.Now()
		user.Active = true
		return nil
	}
}

func (m *UserModel) Authenticate(email, password string) (uint, error) {
	switch email {
	case mockUser.Email:
		return 1, nil
	default:
		return 0, models.ErrInvalidCredentials
	}
}

func (m *UserModel) Get(id uint) (*models.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, models.ErrNoRecord
	}
}
