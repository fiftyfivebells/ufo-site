package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no existing record found")

type Sighting struct {
	Index     int
	UserID    int
	Datetime  time.Time
	Season    string
	City      string
	State     string
	Country   string
	Shape     string
	Duration  int
	Latitude  float64
	Longitude float64
}
