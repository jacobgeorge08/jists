package main

import "github.com/jacobgeorge08/jists/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
