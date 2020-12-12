package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) Routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequests, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Get("/report", dynamicMiddleware.ThenFunc(app.reportSightingForm))
	mux.Post("/report", dynamicMiddleware.ThenFunc(app.reportSighting))
	mux.Get("/stats", dynamicMiddleware.ThenFunc(app.showStatistics))
	mux.Get("/sighting/:id", dynamicMiddleware.ThenFunc(app.showSighting))
	mux.Get("/sightings/:state", dynamicMiddleware.ThenFunc(app.showSightings))

	// file server to serve static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
