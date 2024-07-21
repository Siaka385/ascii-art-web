package main

import (
	"log"
	"net/http"

	"ascii-art/asciifunc"
)

func main() {
	// mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", asciifunc.Indexhandler)
	// mux.HandleFunc("/400", asciifunc.BadRequest)

	// log.Println("starting server on: http://localhost:8080")
	// log.Fatal(http.ListenAndServe("localhost:8080", mux))
	http.HandleFunc("/asciihandler", asciifunc.Trial)
	http.HandleFunc("/400", asciifunc.BadRequest)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8084", nil))
}
