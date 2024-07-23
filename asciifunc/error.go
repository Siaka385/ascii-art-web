package asciifunc

import (
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

func StatusInternalServerError(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("500.html")
	if err != nil {
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
		return
	}
	if r.URL.Query().Get("error") == "true" {
		err = tmp.Execute(w, nil)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("400.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if r.URL.Query().Get("error") == "true" {
		tmpl.Execute(w, nil)
	}
}

func LoadContentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusBadRequest)
}

func LoadContentHandlerInternalError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusInternalServerError)
}

func Pagenofound(w http.ResponseWriter) {
	RenderTemplate(w, "404.html", nil, http.StatusNotFound)
}
