package main

import "github.com/bueti/github-analyzer/internal/models"

// templateData holds any dynamic data that we want to pass to our HTML templates.
type templateData struct {
	Repo *models.Repo
}
