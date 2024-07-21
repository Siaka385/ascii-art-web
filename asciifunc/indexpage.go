package asciifunc

import (
	"log"
	"net/http"
	"text/template"
)

func Indexhandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("indext.html")
	if err != nil {
		log.Fatalf("ERROR PARSING TEMPLATE %v", err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalf("ERROR EXECUTING TEMPLATE %v", err)
	}
}
