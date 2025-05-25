package main

import "snippetbox.dmytron.gogo/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
