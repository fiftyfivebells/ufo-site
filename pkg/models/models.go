package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no existing record found")

type Sighting struct {
	index     int
	userID    int
	datetime  time.Time
	season    string
	city      string
	state     string
	country   string
	shape     string
	duration  int
	latitude  float64
	longitude float64
}
