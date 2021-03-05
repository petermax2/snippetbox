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

	tx := m.DB.Limit(1).Find(&snippet, id)
	if tx.RowsAffected < 1 && tx.Error == nil {
		return nil, ErrNoRecord
	}
	return snippet, tx.Error
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	var snippets []*Snippet
	tx := m.DB.Limit(10).Order("id DESC").Find(&snippets, "expires > CURRENT_TIMESTAMP")
	return snippets, tx.Error
}
