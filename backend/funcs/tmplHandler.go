package groupie

import (
	"fmt"
	"html/template"
	"net/http"
)

func TmplHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		HandleErrorPage(w, http.StatusNotFound)
		return
	}

	artists := ApiArtists()
	locations := ApiLocations()
	dates := ApiDates()
	relations := ApiRelation()

	data := PageData{
		Artists:   artists,
		Locations: locations,
		Dates:     dates,
		Relations: relations,
	}

	tmpl, err := template.ParseFiles("../../frontend/index.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing template: %v", err), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		return
	}
}
