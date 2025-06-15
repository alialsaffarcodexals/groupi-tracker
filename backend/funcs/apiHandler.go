package groupie

import (
	"encoding/json"
	"fmt"
)

func ApiMain() []string {
	apiURL := "https://groupietrackers.herokuapp.com/api"
	body := ReadUrl(apiURL)

	var main Main
	err := json.Unmarshal(body, &main)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON from %s: %v\n", apiURL, err)
		return nil
	}
	arr := []string{main.Artists, main.Locations, main.Dates, main.Relation}
	return arr
}

func ApiArtists() []Artists {
	mainArr := ApiMain()
	apiURL := mainArr[0]
	body := ReadUrl(apiURL)

	var artists []Artists
	err := json.Unmarshal(body, &artists)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON from %s: %v\n", apiURL, err)
		return nil
	}
	return artists
}

func ApiLocations() []Locations {
	mainArr := ApiMain()
	apiURL := mainArr[1]
	body := ReadUrl(apiURL)

	var index LocationsIndex
	err := json.Unmarshal(body, &index)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON from %s: %v\n", apiURL, err)
		return nil
	}
	locations := index.Locations
	return locations
}

func ApiDates() []Dates {
	mainArr := ApiMain()
	apiURL := mainArr[2]
	body := ReadUrl(apiURL)

	var index DatesIndex
	err := json.Unmarshal(body, &index)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON from %s: %v\n", apiURL, err)
		return nil
	}
	dates := index.Dates
	return dates
}

func ApiRelation() []Relations {
	mainArr := ApiMain()
	apiURL := mainArr[3]
	body := ReadUrl(apiURL)

	var index RelationsIndex
	err := json.Unmarshal(body, &index)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON from %s: %v\n", apiURL, err)
		return nil
	}
	relations := index.Relation
	return relations
}
