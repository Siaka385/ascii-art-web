package main

import (
	"log"
	"net/http"

	"ascii-art/asciifunc"
)

func router(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		asciifunc.Indexhandler(w, r)
	} else if r.URL.Path == "/asciihandler" {
		asciifunc.Trial(w, r)
	} else {
		asciifunc.Pagenofound(w)
		return
	}
}

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", router)
	mux.HandleFunc("/400", func(w http.ResponseWriter, r *http.Request) {
		asciifunc.RenderTemplate(w, "400.html", nil, http.StatusBadRequest)
	})

	log.Println("starting server on: http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8085", mux))
}
