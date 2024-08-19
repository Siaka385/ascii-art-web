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

	//for file server
	mux.HandleFunc("/static/", asciifunc.Fileserver)

	mux.HandleFunc("/", asciifunc.Router)
	// Modify your Go server code to listen on all interfaces (0.0.0.0) instead of just localhost:
	log.Println("starting server on: http://0.0.0.0:8080")
	if err := http.ListenAndServe("0.0.0.0:8088", mux); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
