package groupie

import (
	"net/http"
	"text/template"
)

func HandleErrorPage(w http.ResponseWriter, statusCode int) {
	var file string
	switch statusCode {
	case http.StatusNotFound:
		file = "../../frontend/404.html"
	case http.StatusInternalServerError:
		file = "../../frontend/500.html"
	default:
		file = "../../frontend/404.html"
	}

	tmpl, err := template.ParseFiles(file)
	if err != nil {
		http.Error(w, "Error page not found", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	tmpl.Execute(w, nil)
}
