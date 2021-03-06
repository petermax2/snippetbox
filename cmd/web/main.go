package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golangcollege/sessions"
	"nirpet.at/snippetbox/pkg/models"
)

// this struct holds application wide dependencies
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	dbDSN := flag.String("dsn", DEFAULT_DSN, "DSN of the PostgreSQL database to connect to")
	serverAddress := flag.String("address", ":8080", "Network address (and port) of the Snippetbox web server")
	secret := flag.String("secret", "aishoifee*r?ekuk7Mee9Rahhu3juh/i", "Secret key to use for session management")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO]  ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

	templateCache, err := newTemplateCache("./ui/html")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 8 * time.Hour

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		templateCache: templateCache,
	}

	app.initModels(openDB(*dbDSN, errorLog))

	server := &http.Server{
		Addr:     *serverAddress,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting web server on %s", *serverAddress)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}
