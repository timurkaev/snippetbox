package main

import "github.com/timurkaev/snippetbox/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
