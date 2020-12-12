package main

import (
	"html/template"
	"net/url"
	"path/filepath"
	"stephenbell.dev/ufo-site/pkg/models"
)

type templateData struct {
	Flash       string
	Sighting    *models.Sighting
	Sightings   []*models.Sighting
	CurrentYear int
	FormData    url.Values
	FormErrors  map[string]string
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.tmpl"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.tmpl"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.tmpl"))
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
