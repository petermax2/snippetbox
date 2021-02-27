package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// app configuration - cmd line arguments + environment variables (if any)
	serverAddress := flag.String("address", ":8080", "Network address (and port) of the server")
	flag.Parse()

	// logging
	infoLog := log.New(os.Stdout, "[INFO]  ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

	// HTTP message routing
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// file server for static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// start the web server
	server := &http.Server{
		Addr:     *serverAddress,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *serverAddress)
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}
