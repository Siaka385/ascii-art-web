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

	img := http.FileServer(http.Dir("images"))
	mux.Handle("/images/", http.StripPrefix("/images/", img))

	mux.HandleFunc("/", asciifunc.Router)
	mux.HandleFunc("/asciihandler", asciifunc.Trial)
	mux.HandleFunc("/400", asciifunc.BadRequest)
	mux.HandleFunc("/400/badrequest", asciifunc.LoadContentHandler)
	mux.HandleFunc("/500", asciifunc.StatusInternalServerError)
	mux.HandleFunc("/internalservererror", asciifunc.LoadContentHandlerInternalError)
	mux.HandleFunc("/405", asciifunc.Wrongmethodused)

	log.Println("starting server on: http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8085", mux))
}
