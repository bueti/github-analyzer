package main

import (
	"html/template"
	"path/filepath"

	"github.com/bueti/github-analyzer/internal/models"
)

// templateData holds any dynamic data that we want to pass to our HTML templates.
type templateData struct {
	Repo *models.Repo
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.go.html")
	if err != nil {
		return nil, err
	}
	pages = append(pages, "./ui/html/base.go.html", "./ui/html/partials/nav.go.html")

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles(pages...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
