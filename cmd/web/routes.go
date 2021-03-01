package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// HTTP message routing
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/api/snippet", app.apiGetSnippet)
	mux.HandleFunc("/api/snippet/create", app.apiCreateSnippet)
	mux.HandleFunc("/api/snippet/latest", app.apiGetLatestSnippets)

	// file server for static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
