package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"nirpet.at/snippetbox/pkg/forms"
	"nirpet.at/snippetbox/pkg/models"
)

func (app *application) htmlHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.renderHtml(w, r, "home.page.tmpl", &templateData{Snippets: snippets})
}

func (app *application) htmlShowSnippet(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id, err := strconv.Atoi(urlParams["id"])
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if err == models.ErrNoRecord {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.renderHtml(w, r, "show.page.tmpl", &templateData{Snippet: snippet})
}

func (app *application) htmlCreateSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.renderHtml(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) htmlCreateSnippet(w http.ResponseWriter, r *http.Request) {
	// parse HTML form data
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// user input validation
	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValues("expires", "365", "7", "1")

	if !form.Valid() {
		app.renderHtml(w, r, "create.page.tmpl", &templateData{Form: form})
		return
	}

	// convert expires to integer
	expires, err := strconv.Atoi(form.Get("expires"))
	if err != nil {
		app.serverError(w, err)
		return
	}

	// create new snippet from form data
	snippet := &models.Snippet{
		Title:   form.Get("title"),
		Content: form.Get("content"),
		Expires: time.Now().AddDate(0, 0, expires),
	}

	// DB interaction - save snippet via ORM
	err = app.snippets.Insert(snippet)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// display created snippet to the user
	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", snippet.ID), http.StatusSeeOther)
}

func (app *application) apiGetSnippet(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	id, err := strconv.Atoi(urlParams["id"])
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		app.notFound(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snippet)
}

func (app *application) apiGetLatestSnippets(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.clientError(w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snippets)
}

func (app *application) apiCreateSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		app.clientError(w, http.StatusUnsupportedMediaType)
		return
	}

	var snippet models.Snippet
	err := json.NewDecoder(r.Body).Decode(&snippet)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.snippets.Insert(&snippet)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(snippet)
}
