package models

import (
	"errors"
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

func (m *UserModel) Authenticate(email, password string) (uint, error) {
	user := &User{}

	tx := m.DB.Limit(1).Where("email = ?", email).Find(&user)
	if tx.RowsAffected < 1 && tx.Error == nil {
		return 0, ErrInvalidCredentials
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return user.ID, nil
}

func (m *UserModel) Get(id int) (*User, error) {
	return nil, nil
}
