package main

import (
	"net/http"

	"github.com/bueti/github-analyzer/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	app.render(w, http.StatusOK, "base.go.html", &templateData{
		Repo: &models.Repo{},
	})
}

func (app *application) repoView(w http.ResponseWriter, r *http.Request) {
	org := r.URL.Query().Get("org")
	if org == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	repo := r.URL.Query().Get("repo")
	if repo == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	repoInfo, err := app.repos.Get(org, repo)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, http.StatusOK, "view.go.html", &templateData{
		Repo: repoInfo,
	})
}
