package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, http.StatusOK, "home.go.html", data)
}

func (app *application) repoView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	if params.ByName("org") == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	if params.ByName("repo") == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	repoInfo, err := app.repos.Get(params.ByName("org"), params.ByName("repo"))
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Repo = repoInfo

	app.render(w, http.StatusOK, "view.go.html", data)
}
