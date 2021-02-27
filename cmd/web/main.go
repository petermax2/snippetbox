package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	serverAddress := flag.String("address", ":8080", "Network address (and port) of the server")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on %s", *serverAddress)
	err := http.ListenAndServe(*serverAddress, mux)
	log.Fatal(err)
}
