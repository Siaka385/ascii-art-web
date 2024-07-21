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

func StatusInternalServerError(w http.ResponseWriter) {
	RenderTemplate(w, "500.html", nil, http.StatusInternalServerError)
}

func Badrequest(w http.ResponseWriter, data string) {
	RenderTemplate(w, "asciipage.html", data, http.StatusBadRequest)
}

func Pagenofound(w http.ResponseWriter) {
	RenderTemplate(w, "404.html", nil, http.StatusNotFound)
}
