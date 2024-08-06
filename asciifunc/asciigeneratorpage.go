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
		CheckError(w, r)
	} else if r.URL.Path == "/400" {
		BadRequest(w, r)
		return
	} else if r.URL.Path == "/400/badrequest" {
		Setstatus(w, r, http.StatusBadRequest)
		return
	} else if r.URL.Path == "/500" {
		StatusInternalServerError(w, r)
		return
	} else if r.URL.Path == "/internalservererror" {
		Setstatus(w, r, http.StatusInternalServerError)
		return
	} else if r.URL.Path == "/405" {
		Wrongmethodused(w, r)
		return
	} else if r.URL.Path == "/404banner" {
		StatusUnavailableBanner(w, r)
		return
	} else if r.URL.Path == "/unavailablebanner" {
		Setstatus(w, r, http.StatusNotFound)
		return
	} else {
		PageNotFound(w, "404.html", nil, http.StatusNotFound)
		return
	}
}

func CheckError(w http.ResponseWriter, r *http.Request) {
	method := strings.ToUpper(r.Method)

	if method != "POST" {

		log.Println("wrong http method encountered, redirecting to /405 page")
		http.Redirect(w, r, "/405", http.StatusFound)
		return
	}
	err := r.ParseForm()
	if err != nil {
		log.Println("Internal server error encountered, redirecting to /500 page")
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}

	if !IsBannerAccepted(r.Form.Get("Banner")) {
		log.Println("Bannner unavailabe redirecting to page 404banner")
		http.Redirect(w, r, "/404banner", http.StatusFound)
		return
	}

	if !CheckFileAuthencity(r.Form.Get("Banner")) {
		log.Println("Internal server error encountered, redirecting to /500 page")
		http.Redirect(w, r, "/500?error=true", http.StatusFound)
		return
	}

	if !IsItAnAsciiCharacter(r.Form.Get("input-text")) {
		log.Println("Non-ASCII character found, redirecting to /400")
		http.Redirect(w, r, "/400", http.StatusFound)

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
		log.Println("Internal server error encountered, redirecting to /500 page")
		http.Redirect(w, r, "/500?error=true", http.StatusFound)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Internal server error encountered, redirecting to /500 page")
		http.Redirect(w, r, "/500?error=true", http.StatusFound)
		return
	}
}
