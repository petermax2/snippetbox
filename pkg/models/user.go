package models

import "gorm.io/gorm"

type UserModel struct {
	DB *gorm.DB
}

func (m *UserModel) Migrate() error {
	return m.DB.AutoMigrate(&User{})
}

func (m *UserModel) Insert(user *User) error {
	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	return -1, nil
}

func (m *UserModel) Get(id int) (*User, error) {
	return nil, nil
}
