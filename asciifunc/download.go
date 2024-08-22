package asciifunc

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Download(w http.ResponseWriter, r *http.Request) {
	fileContents, err := os.ReadFile("tester.txt")

	if len(fileContents) == 0 {
		return
	}

	if err != nil {
		log.Println("Internal server error encountered, redirecting to /500 page")
		http.Redirect(w, r, "/500?error=true", http.StatusFound)
		return
	}

	fileinfo, er := os.Stat("tester.txt")

	if er != nil {
		log.Println("Internal server error encountered, redirecting to /500 page")
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}

	// Set headers to suggest a filename and prompt a download
	w.Header().Set("Content-Disposition", "attachment; filename=asciipage.txt")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileinfo.Size()))
	w.Header().Set("Content-Type", "text/plain")
	os.Truncate("tester.txt", 0)
	_, err = w.Write(fileContents)
	if err != nil {
		log.Println("Internal server error encountered, redirecting to /500 page")
		http.Redirect(w, r, "/500", http.StatusFound)
		return
	}

	log.Print("Downloading")
}
