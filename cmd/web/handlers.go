package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"stephenbell.dev/ufo-site/pkg/models"
)

// Route handler for the home page
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.renderTemplate(w, r, "home.page.tmpl", nil)
}

// Route handler for the form for reporting a sighting
func (app *application) reportSightingForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Report a sighting..."))
}

// Route handler for the sighting reporter
func (app *application) reportSighting(w http.ResponseWriter, r *http.Request) {

	userID := 10
	datetime := time.Now()
	season := "fall"
	city := "boston"
	state := "ma"
	country := "us"
	shape := "triangle"
	duration := 180
	lat := 71.01040
	long := -43.0220

	id, err := app.sightings.Insert(userID, datetime, season, city, state, country, shape, duration, lat, long)

	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/sighting/:%d", id), http.StatusSeeOther)
}

// Route handler for the statistics page
func (app *application) showStatistics(w http.ResponseWriter, r *http.Request) {
	return
}

// Route handler for sightings page
func (app *application) showSightings(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get(":state")

	s, err := app.sightings.GetByState(state)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	app.renderTemplate(w, r, "sightings.page.tmpl", &templateData{Sightings: s})
}

// Route handler for showing an individual sighting
func (app *application) showSighting(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	s, err := app.sightings.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.renderTemplate(w, r, "show.page.tmpl", &templateData{Sighting: s})
}
