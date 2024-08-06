package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"ascii-art/asciifunc"
)

func main() {
	myargs := os.Args[1:]
	if len(myargs) > 0 {
		fmt.Println("Too many args used in the terminal")
		os.Exit(1)
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", asciifunc.Router)

	log.Println("starting server on: http://localhost:8086")
	log.Fatal(http.ListenAndServe("localhost:8087", mux))
}
