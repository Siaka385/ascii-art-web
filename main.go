package main

import (
	"log"
	"net/http"

	"ascii-art/asciifunc"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", asciifunc.Router)

	log.Println("starting server on: http://localhost:8087")
	log.Fatal(http.ListenAndServe("localhost:8086", mux))
}
