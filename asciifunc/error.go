package asciifunc

import (
	"log"
	"net/http"
	"text/template"
)

func PageNotFound(w http.ResponseWriter, tmpl string, data interface{}, statusCode int) {
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

	err = tmp.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func StatusUnavailableBanner(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("banner404.html")
	if err != nil {
		log.Println("Internal server error encountered, redirecting to /500 page")
		http.Redirect(w, r, "/500?error=true", http.StatusFound)
		return
	}

	err = tmp.Execute(w, nil)
	if err != nil {
		log.Println("Internal server error encountered, redirecting to /500 page")
		http.Redirect(w, r, "/500?error=true", http.StatusFound)
		return
	}
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("400.html")
	if err != nil {
		log.Println("Internal server error encountered, redirecting to /500 page")
		http.Redirect(w, r, "/500?error=true", http.StatusFound)
		return
	}
	tmpl.Execute(w, nil)
}

func Setstatus(w http.ResponseWriter, r *http.Request, statusCode int) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(statusCode)
}

func Wrongmethodused(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("405.html")
	if err != nil {
		log.Println("Internal server error encountered, redirecting to /500 page")
		http.Redirect(w, r, "/500?error=true", http.StatusFound)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	err = tmp.Execute(w, nil)
	if err != nil {
		log.Println("Internal server error encountered, redirecting to /500 page")
		http.Redirect(w, r, "/500?error=true", http.StatusFound)
		return
	}
}
