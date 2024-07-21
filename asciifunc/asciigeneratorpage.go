package asciifunc

import (
	"net/http"
	"strings"
	"text/template"
)

type Data struct {
	Result string
}

func Trial(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if !CheckFileEmpty(r.Form.Get("Banner")) || !CheckNumberOfLinesInTheFile(r.Form.Get("Banner")) {
		StatusInternalServerError(w)
		return
	}
	if !IsItAnAsciiCharacter(r.Form.Get("input-text")) {
		data := Data{
			Result: "Bad request",
		}
		Badrequest(w, data.Result)
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
