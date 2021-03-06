package models

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserModel struct {
	DB *gorm.DB
}

func (m *UserModel) Migrate() error {
	return m.DB.AutoMigrate(&User{})
}

func (m *UserModel) Insert(user *User) error {
	// use bcrypt to store the password in hashed form
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 13)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)

	tx := m.DB.Create(user)
	if strings.Contains(tx.Error.Error(), "duplicate key") && strings.Contains(tx.Error.Error(), "email") {
		return ErrDuplicateEmail
	}
	return tx.Error
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	return -1, nil
}

func (m *UserModel) Get(id int) (*User, error) {
	return nil, nil
}
