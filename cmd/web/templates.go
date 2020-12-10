package main

import "stephenbell.dev/ufo-site/pkg/models"

type templateData struct {
	Sighting  *models.Sighting
	Sightings []*models.Sighting
}
