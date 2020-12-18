package models

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no existing record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Sighting struct {
	Index     int            `json:"index"`
	UserID    int            `json:"user_id"`
	Datetime  time.Time      `json:"datetime"`
	Season    string         `json:"season"`
	City      string         `json:"city"`
	State     string         `json:"state"`
	Country   string         `json:"country"`
	Shape     sql.NullString `json:"shape"`
	Duration  sql.NullInt64  `json:"duration"`
	Latitude  float64        `json:"lat"`
	Longitude float64        `json:"long"`
	Sighted   int            `json:"sighted"`
}

type User struct {
	ID         int
	Username   string
	Email      string
	HashedPass []byte
	Created    time.Time
	Active     bool
}
