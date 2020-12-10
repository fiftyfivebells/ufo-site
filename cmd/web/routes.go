package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) Routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequests, secureHeaders)

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/report", http.HandlerFunc(app.reportSightingForm))
	mux.Post("/report", http.HandlerFunc(app.reportSighting))
	mux.Get("/stats", http.HandlerFunc(app.showStatistics))
	mux.Get("/sighting/:id", http.HandlerFunc(app.showSighting))
	mux.Get("/sightings/:state", http.HandlerFunc(app.showSightings))

	// file server to serve static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
