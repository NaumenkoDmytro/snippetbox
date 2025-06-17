package mocks

import (
	"time"

	"snippetbox.dmytron.gogo/internal/models"
)

var mockSnippet = models.Snippet{
	ID:      1,
	Title:   "An old silent pond",
	Content: "An old silent pond...",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnipppetModel struct{}

func (m *SnipppetModel) Insert(title string, content string, expires int) (int, error) {
	return 2, nil
}

func (m *SnipppetModel) Get(id int) (models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return models.Snippet{}, models.ErrNoRecord
	}
}

func (m *SnipppetModel) Latest() ([]models.Snippet, error) {
	return []models.Snippet{mockSnippet}, nil
}
