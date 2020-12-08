package main

import "net/http"

func (app *application) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/report", app.reportSighting)
	mux.HandleFunc("/stats", app.showStatistics)
	mux.HandleFunc("/sightings", app.showSightings)

	// file server to serve static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return secureHeaders(mux)
}