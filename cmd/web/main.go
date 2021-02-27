package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// this struct holds application wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// app configuration - cmd line arguments + environment variables (if any)
	serverAddress := flag.String("address", ":8080", "Network address (and port) of the server")
	flag.Parse()

	// logging
	infoLog := log.New(os.Stdout, "[INFO]  ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

	// initialize the application instance
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// HTTP message routing
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

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
