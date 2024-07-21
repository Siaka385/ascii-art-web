package asciifunc

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Data struct {
	Result string
}

func Router(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		Indexhandler(w, r)
	} else if r.URL.Path == "/asciihandler" {
		Trial(w, r)
	} else if r.URL.Path == "/400" {
		BadRequest(w, r)
		return
	} else {
		Pagenofound(w)
		return
	}
}

func Trial(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if !IsItAnAsciiCharacter(r.Form.Get("input-text")) {
		log.Println("Non-ASCII character found, responding with 400")
		data:=Data{
			Result: "AN NON ASCII CHARACTER USED",
		}
        w.Header().Set("Content-Type", "text/html")
        w.WriteHeader(http.StatusOK)
        tmpl, _ := template.ParseFiles("400.html")
        tmpl.Execute(w, data)
        return
	}
	Asciihandler(w, r)
}
func Asciihandler(w http.ResponseWriter, r *http.Request) {
	var print string

	if strings.Contains(r.Form.Get("input-text"), "\n") {
		myslice := strings.Split(r.Form.Get("input-text"), "\n")

		for i := 0; i < len(myslice); i++ {
			print += Asciigenerate([]string{myslice[i], r.Form.Get("Banner")})
		}

	} else {
		print += Asciigenerate([]string{r.Form.Get("input-text"), r.Form.Get("Banner")})
	}

	data := Data{
		Result: print,
	}

	tmpl, err := template.ParseFiles("asciipage.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}