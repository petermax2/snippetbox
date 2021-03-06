package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	defaultMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// HTTP message routing
	router := mux.NewRouter().StrictSlash(true)

	// register HTML routes
	router.HandleFunc("/", app.htmlHome).Methods("GET")
	router.HandleFunc("/snippet/{id}", app.htmlShowSnippet).Methods("GET")
	router.HandleFunc("/snippet/create", app.htmlCreateSnippetForm).Methods("GET")
	router.HandleFunc("/snippet/create", app.htmlCreateSnippet).Methods("POST")

	// register API routes (JSON)
	router.HandleFunc("/api/snippet/{id}", app.apiGetSnippet).Methods("GET")
	router.HandleFunc("/api/snippet/create", app.apiCreateSnippet).Methods("POST")
	router.HandleFunc("/api/snippet/latest", app.apiGetLatestSnippets).Methods("GET")

	// file server for static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handle("/static/", http.StripPrefix("/static", fileServer))

	return defaultMiddleware.Then(router)
}
