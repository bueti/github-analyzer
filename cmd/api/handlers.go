package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./ui/html/base.go.html",
		"./ui/html/partials/nav.go.html",
		"./ui/html/pages/home.go.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
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

	fmt.Fprintf(w, "Display a specific repo (org: %s, repo: %s)", org, repo)
}
