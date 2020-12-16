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
	Index     int
	UserID    int
	Datetime  time.Time
	Season    string
	City      string
	State     string
	Country   string
	Shape     sql.NullString
	Duration  sql.NullInt64
	Latitude  float64
	Longitude float64
	Sighted   int
}

type User struct {
	ID         int
	Username   string
	Email      string
	HashedPass []byte
	Created    time.Time
	Active     bool
}
