package main

import (
	"html/template"
	"log"
	"net/http"
)

// Route handler for the home page
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", 500)
	}
}

// Route handler for the sighting reporter
func reportSighting(w http.ResponseWriter, r *http.Request) {
	return
}

// Route handler for the statistics page
func showStatistics(w http.ResponseWriter, r *http.Request) {
	return
}

// Route handler for sightings page
func showSightings(w http.ResponseWriter, r *http.Request) {
	return
}
