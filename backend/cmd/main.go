package main

import (
	"fmt"
	groupie "groupie/funcs"
	"net/http"
)

func main() {
	http.HandleFunc("/", groupie.TmplHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../frontend/static"))))
	http.HandleFunc("/artist/", groupie.ArtistDetailsHandler)

	fmt.Print("Server started at ")
	fmt.Println("\033[32mhttp://localhost:8080\033[0m")

	http.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		groupie.HandleErrorPage(w, http.StatusNotFound)
	})

	http.HandleFunc("/500", func(w http.ResponseWriter, r *http.Request) {
		groupie.HandleErrorPage(w, http.StatusInternalServerError)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Server error: " + err.Error())
	}
}
