package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) Routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequests, secureHeaders)

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/report", app.reportSighting)
	mux.HandleFunc("/stats", app.showStatistics)
	mux.HandleFunc("/sighting", app.showSighting)
	mux.HandleFunc("/sightings", app.showSightings)

	// file server to serve static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
