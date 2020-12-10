package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"stephenbell.dev/ufo-site/pkg/models"
)

// Route handler for the home page
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

// Route handler for the sighting reporter
func (app *application) reportSighting(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
	}

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

	http.Redirect(w, r, fmt.Sprintf("/sighting?id=%d", id), http.StatusSeeOther)
}

// Route handler for the statistics page
func (app *application) showStatistics(w http.ResponseWriter, r *http.Request) {
	return
}

// Route handler for sightings page
func (app *application) showSightings(w http.ResponseWriter, r *http.Request) {
	return
}

// Route handler for showing an individual sighting
func (app *application) showSighting(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
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

	data := &templateData{Sighting: s}

	files := []string{
		"./ui/html/show.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}
}
