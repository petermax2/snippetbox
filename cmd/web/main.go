package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx"
)

// this struct holds application wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	serverAddress := flag.String("address", ":8080", "Network address (and port) of the server")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO]  ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog := log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	dbConfig := &pgx.ConnConfig{
		Host:     "localhost",
		Port:     8082,
		Database: "snippetbox",
		User:     "web",
		Password: "password",
	}

	conn, err := pgx.Connect(*dbConfig)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer conn.Close()

	server := &http.Server{
		Addr:     *serverAddress,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *serverAddress)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}
