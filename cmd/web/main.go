package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":5000", "HTTP network address")

	flag.Parse()

	// Create new ServeMux
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Create file server for static assets
	fileServer := http.FileServer(http.Dir("./assets/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
