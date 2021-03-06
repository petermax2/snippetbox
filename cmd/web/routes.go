package main

import "net/http"

func (app *application) routes() http.Handler {
	// HTTP message routing
	mux := http.NewServeMux()

	// register HTML routes
	mux.HandleFunc("/", app.htmlShowHome)
	mux.HandleFunc("/snippet", app.htmlShowSnippet)

	// register API routes (JSON)
	mux.HandleFunc("/api/snippet", app.apiGetSnippet)
	mux.HandleFunc("/api/snippet/create", app.apiCreateSnippet)
	mux.HandleFunc("/api/snippet/latest", app.apiGetLatestSnippets)

	// file server for static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return app.logRequest(secureHeaders(mux))
}
