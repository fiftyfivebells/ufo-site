package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"stephenbell.dev/ufo-site/pkg/forms"
	"stephenbell.dev/ufo-site/pkg/models"
)

// Route handler for the home page
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.renderTemplate(w, r, "home.page.tmpl", nil)
}

// Route handler for the form for reporting a sighting
func (app *application) reportSightingForm(w http.ResponseWriter, r *http.Request) {
	app.renderTemplate(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

// Route handler for the sighting reporter
func (app *application) reportSighting(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)

	form.Required("city", "state", "shape", "duration")

	if !form.Valid() {
		app.renderTemplate(w, r, "create.page.tmpl", &templateData{Form: form})
		return
	}

	dur, err := strconv.Atoi(form.Get("duration"))
	if err != nil {
		app.serverError(w, err)
		return
	}

	lat, long := app.getLatAndLong(form.Get("city"), form.Get("state"))
	state := strings.ToLower(form.Get("state"))
	state = app.getStateAbbrev(state)
	season := app.convertTimeToSeason(time.Now())

	id, err := app.sightings.Insert(0, time.Now(), season, form.Get("city"), state, "us", form.Get("shape"), dur, lat, long)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.session.Put(r, "flash", "Sighting successfully reported!")
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

	app.renderTemplate(w, r, "show.page.tmpl", &templateData{
		Sighting: s,
	})
}

// Display the register user form
func (app *application) registerUserForm(w http.ResponseWriter, r *http.Request) {
	app.renderTemplate(w, r, "register.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

// Insert user into the database
func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Register the user...")
}

// Display the user login form
func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display the user login form...")
}

// Log the user into the site
func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Log the user in...")
}

// Log the user out
func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Log the user out...")
}
