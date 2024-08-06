package asciifunc

import (
	"log"
	"net/http"
	"text/template"
)

func Indexhandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("indext.html")
	if err != nil {
		log.Println("Internal server error encountered, redirecting to /404 page")
		http.Redirect(w, r, "/notfound", http.StatusFound)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Internal server error encountered, redirecting to /404 page")
		http.Redirect(w, r, "/notfound", http.StatusFound)
		return
	}
}
