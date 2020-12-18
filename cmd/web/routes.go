package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) Routes() http.Handler {
	standardMiddleware := alice.New(app.recoverPanic, app.logRequests, secureHeaders)
	dynamicMiddleware := alice.New(app.session.Enable, noSurf, app.authenticate)

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.home))
	mux.Post("/", dynamicMiddleware.ThenFunc(app.getPrediction))
	mux.Get("/report", dynamicMiddleware.ThenFunc(app.reportSightingForm))
	mux.Post("/report", dynamicMiddleware.Append(app.requireAuthorization).ThenFunc(app.reportSighting))
	mux.Get("/stats", dynamicMiddleware.ThenFunc(app.showStatistics))
	mux.Get("/sightings", dynamicMiddleware.ThenFunc(app.showSightings))
	mux.Get("/sightings/all", dynamicMiddleware.ThenFunc(app.sendSightings))
	mux.Get("/sighting/:id", dynamicMiddleware.ThenFunc(app.showSighting))
	mux.Get("/sightings/:state", dynamicMiddleware.ThenFunc(app.showSightings))

	// User registrationg, login, and logout
	mux.Get("/user/register", dynamicMiddleware.ThenFunc(app.registerUserForm))
	mux.Post("/user/register", dynamicMiddleware.ThenFunc(app.registerUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthorization).ThenFunc(app.logoutUser))

	// file server to serve static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
