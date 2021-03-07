package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	defaultMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable, noSurf, app.authenticate)
	authenticatedMiddleware := dynamicMiddleware.Append(app.requireAuthentication)

	// HTTP message routing
	router := mux.NewRouter()

	router.Handle("/", dynamicMiddleware.ThenFunc(app.htmlHome)).Methods("GET")
	router.HandleFunc("/ping", ping).Methods("GET")

	// register HTML routes (snippet)
	router.Handle("/snippet/create", authenticatedMiddleware.ThenFunc(app.htmlCreateSnippetForm)).Methods("GET")
	router.Handle("/snippet/create", authenticatedMiddleware.ThenFunc(app.htmlCreateSnippet)).Methods("POST")
	router.Handle("/snippet/{id}", dynamicMiddleware.ThenFunc(app.htmlShowSnippet)).Methods("GET")

	// register HTML routes (user)
	router.Handle("/user/signup", dynamicMiddleware.ThenFunc(app.htmlSignupUserForm)).Methods("GET")
	router.Handle("/user/signup", dynamicMiddleware.ThenFunc(app.htmlSignupUser)).Methods("POST")
	router.Handle("/user/login", dynamicMiddleware.ThenFunc(app.htmlLoginUserForm)).Methods("GET")
	router.Handle("/user/login", dynamicMiddleware.ThenFunc(app.htmlLoginUser)).Methods("POST")
	router.Handle("/user/logout", authenticatedMiddleware.ThenFunc(app.htmlLogoutUser)).Methods("POST")

	// register API routes (JSON)
	// NOTES these routes have nothing to do with the book
	//       I keep them here as a reminder for later projects
	router.HandleFunc("/api/snippet/create", app.apiCreateSnippet).Methods("POST")
	router.HandleFunc("/api/snippet/latest", app.apiGetLatestSnippets).Methods("GET")
	router.HandleFunc("/api/snippet/{id}", app.apiGetSnippet).Methods("GET")

	// file server for static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.PathPrefix("/static").Handler(http.StripPrefix("/static/", fileServer))

	return defaultMiddleware.Then(router)
}
