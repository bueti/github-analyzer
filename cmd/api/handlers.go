package main

import (
	"errors"
	"net/http"

	"github.com/bueti/github-analyzer/internal/models"
	"github.com/bueti/github-analyzer/internal/validator"
)

type repoViewForm struct {
	Org  string
	Repo string
	validator.Validator
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = repoViewForm{}
	app.render(w, http.StatusOK, "home.go.html", data)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "about.go.html", data)
}

func (app *application) repoView(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := repoViewForm{
		Org:  r.PostForm.Get("org"),
		Repo: r.PostForm.Get("repo"),
	}

	form.CheckField(validator.NotBlank(form.Org), "org", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Org, 39), "org", "This field cannot be more than 39 characters long")
	form.CheckField(validator.NotBlank(form.Repo), "repo", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Repo, 100), "repo", "This field cannot be more than 100 characters long")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "home.go.html", data)
		return
	}

	repoInfo, err := app.repos.Get(form.Org, form.Repo)
	if err != nil {
		if errors.Is(err, models.RepoNotFound) {
			data := app.newTemplateData(r)
			data.Form = form
			app.render(w, http.StatusOK, "home.go.html", data)
			return
		}
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Repo = repoInfo

	app.render(w, http.StatusOK, "view.go.html", data)
}
