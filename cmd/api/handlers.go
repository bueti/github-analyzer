package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/charmbracelet/log"
)

func home(w http.ResponseWriter, r *http.Request) {
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
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func repoView(w http.ResponseWriter, r *http.Request) {
	org := r.URL.Query().Get("org")
	if org == "" {
		log.Error("organization name is not set")
		http.Error(w, "organization name is not set", http.StatusBadRequest)
		return
	}
	repo := r.URL.Query().Get("repo")
	if repo == "" {
		log.Error("repository name is not set")
		http.Error(w, "repository name is not set", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Display a specific repo (org: %s, repo: %s)", org, repo)
}
