package models

import (
	"gorm.io/gorm"
)

type SnippetModel struct {
	DB *gorm.DB
}

func (m *SnippetModel) Migrate() error {
	return m.DB.AutoMigrate(&Snippet{})
}

func (m *SnippetModel) Insert(snippet *Snippet) error {
	tx := m.DB.Create(snippet)
	return tx.Error
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	snippet := &Snippet{}
	tx := m.DB.First(&snippet, id)
	return snippet, tx.Error
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	var snippets []*Snippet
	tx := m.DB.Limit(10).Order("id DESC").Find(&snippets)
	return snippets, tx.Error
}
