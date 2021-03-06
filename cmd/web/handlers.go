package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gorilla/mux"
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
	app.renderHtml(w, r, "create.page.tmpl", nil)
}

func (app *application) htmlCreateSnippet(w http.ResponseWriter, r *http.Request) {
	// parse HTML form data
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")

	// user input validation
	validationErrors := make(map[string]string)

	if strings.TrimSpace(title) == "" {
		validationErrors["title"] = "This field can not be blank"
	} else if utf8.RuneCountInString(title) > 100 {
		validationErrors["title"] = "This field is too long (maximum is 100 characters)"
	}

	if strings.TrimSpace(content) == "" {
		validationErrors["content"] = "This field can not be blank"
	}

	expiresInDays, err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil || expiresInDays < 1 || expiresInDays > 365 {
		validationErrors["expires"] = "This field must be a numeric value between 1 and 365"
	}

	if len(validationErrors) > 0 {
		app.renderHtml(w, r, "create.page.tmpl", &templateData{
			FormErrors: validationErrors,
			FormData:   r.PostForm,
		})
		return
	}

	snippet := &models.Snippet{
		Title:   r.PostForm.Get("title"),
		Content: r.PostForm.Get("content"),
		Expires: time.Now().AddDate(0, 0, expiresInDays),
	}

	err = app.snippets.Insert(snippet)
	if err != nil {
		app.serverError(w, err)
		return
	}

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
