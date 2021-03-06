package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	defaultMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)

	// HTTP message routing
	router := mux.NewRouter()

	// register HTML routes
	router.Handle("/", dynamicMiddleware.ThenFunc(app.htmlHome)).Methods("GET")
	router.Handle("/snippet/create", dynamicMiddleware.ThenFunc(app.htmlCreateSnippetForm)).Methods("GET")
	router.Handle("/snippet/create", dynamicMiddleware.ThenFunc(app.htmlCreateSnippet)).Methods("POST")
	router.Handle("/snippet/{id}", dynamicMiddleware.ThenFunc(app.htmlShowSnippet)).Methods("GET")

	// register API routes (JSON)
	router.HandleFunc("/api/snippet/create", app.apiCreateSnippet).Methods("POST")
	router.HandleFunc("/api/snippet/latest", app.apiGetLatestSnippets).Methods("GET")
	router.HandleFunc("/api/snippet/{id}", app.apiGetSnippet).Methods("GET")

	// file server for static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handle("/static/", http.StripPrefix("/static", fileServer))

	return defaultMiddleware.Then(router)
}
