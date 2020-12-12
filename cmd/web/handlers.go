package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"stephenbell.dev/ufo-site/pkg/models"
)

// Route handler for the home page
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.renderTemplate(w, r, "home.page.tmpl", nil)
}

// Route handler for the form for reporting a sighting
func (app *application) reportSightingForm(w http.ResponseWriter, r *http.Request) {
	app.renderTemplate(w, r, "create.page.tmpl", nil)
}

// Route handler for the sighting reporter
func (app *application) reportSighting(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	city := r.PostForm.Get("city")
	state := r.PostForm.Get("state")
	shape := r.PostForm.Get("shape")
	duration := r.PostForm.Get("duration")

	errors := make(map[string]string)

	if strings.TrimSpace(city) == "" {
		errors["city"] = "This field cannot be blank"
	}

	if strings.TrimSpace(duration) == "" {
		errors["duration"] = "This field cannot be blank"
	}

	if len(errors) > 0 {
		app.renderTemplate(w, r, "create.page.tmpl", &templateData{
			FormErrors: errors,
			FormData:   r.PostForm,
		})
		return
	}

	dur, err := strconv.Atoi(duration)
	if err != nil {
		app.serverError(w, err)
		return
	}

	lat, long := app.getLatAndLong(city, state)
	state = strings.ToLower(state)
	state = app.getStateAbbrev(state)
	season := app.convertTimeToSeason(time.Now())

	id, err := app.sightings.Insert(0, time.Now(), season, city, state, "us", shape, dur, lat, long)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/sighting/%d", id), http.StatusSeeOther)
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
