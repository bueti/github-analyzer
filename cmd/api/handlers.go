package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "home.go.html", data)
}

func (app *application) repoView(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	org := r.PostForm.Get("org")
	repo := r.PostForm.Get("repo")

	repoInfo, err := app.repos.Get(org, repo)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Repo = repoInfo

	app.render(w, http.StatusOK, "view.go.html", data)
}
