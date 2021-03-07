package mock

import (
	"time"

	"gorm.io/gorm"
	"nirpet.at/snippetbox/pkg/models"
)

var mockSnippet = &models.Snippet{
	Model: gorm.Model{
		ID: 1,
	},
	CreatedAt: time.Now(),
	Title:     "Test Snippet",
	Content:   "Unit Testing Stuff",
	Expires:   time.Now().AddDate(0, 0, 3),
}

type SnippetModel struct{}

func (m *SnippetModel) Migrate() error {
	return nil
}

func (m *SnippetModel) Insert(snippet *models.Snippet) error {
	snippet.ID = 2
	snippet.CreatedAt = time.Now()
	snippet.UpdatedAt = time.Now()
	return nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}
