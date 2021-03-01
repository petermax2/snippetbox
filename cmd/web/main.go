package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"nirpet.at/snippetbox/pkg/models"
)

// this struct holds application wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

func main() {
	dbDSN := flag.String("dsn", DEFAULT_DSN, "DSN of the PostgreSQL database to connect to")
	serverAddress := flag.String("address", ":8080", "Network address (and port) of the Snippetbox web server")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO]  ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	app.initModels(openDB(*dbDSN, errorLog))

	server := &http.Server{
		Addr:     *serverAddress,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting web server on %s", *serverAddress)
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}
