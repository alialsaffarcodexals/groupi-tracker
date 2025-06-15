package groupie

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func ArtistDetailsHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 || id > len(ApiArtists()) {
		HandleErrorPage(w, http.StatusNotFound)
		return
	}

	artists := ApiArtists()
	relations := ApiRelation()

	var selected *Artists
	for _, artist := range artists {
		if artist.ID == id {
			selected = &artist
			break
		}
	}

	var concertMap map[string][]string
	for _, rel := range relations {
		if rel.ID == id {
			concertMap = rel.DatesLocations
			break
		}
	}

	data := ArtistDetailPage{
		Artist:   *selected,
		Concerts: concertMap,
	}

	tmpl, err := template.ParseFiles("../../frontend/artist.html")
	if err != nil {
		HandleErrorPage(w, http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, data)
}
