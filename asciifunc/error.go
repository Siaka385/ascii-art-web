package asciifunc

import (
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}, statusCode int) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(statusCode)
	t.Execute(w, data)
}

func StatusInternalServerError(w http.ResponseWriter) {
	RenderTemplate(w, "500.html", nil, http.StatusInternalServerError)
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("BadRequest handler called")

	tmpl, err := template.ParseFiles("400.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusBadRequest)
	tmpl.Execute(w, nil)
}

func Pagenofound(w http.ResponseWriter) {
	RenderTemplate(w, "404.html", nil, http.StatusNotFound)
}
