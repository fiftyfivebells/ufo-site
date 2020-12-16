package postgresql

import (
	"database/sql"
	"errors"
	"time"

	"stephenbell.dev/ufo-site/pkg/models"
)

type SightingModel struct {
	DB *sql.DB
}

func (m *SightingModel) InsertSighting(userID int, datetime time.Time, season, city, state, country, shape string, duration int, lat, long float64, sighted int) (int, error) {

	stmt := `INSERT INTO sightings (user_id, datetime, season, city, state, country, shape, duration, latitude, longitude, sighted) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING index`

	var id int
	err := m.DB.QueryRow(stmt,
		userID,
		datetime,
		season,
		city,
		state,
		country,
		shape,
		duration,
		lat,
		long,
		sighted).Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (m *SightingModel) InsertNoSighting(userID int, datetime time.Time, season, city, state, country string, lat, long float64, sighted int) (int, error) {
	stmt := `INSERT INTO sightings (user_id, datetime, season, city, state, country, latitude, longitude, sighted) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING index`

	var id int
	err := m.DB.QueryRow(stmt,
		userID,
		datetime,
		season,
		city,
		state,
		country,
		lat,
		long,
		sighted).Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (m *SightingModel) Get(id int) (*models.Sighting, error) {
	stmt := `SELECT * FROM sightings WHERE index = $1`
	s := &models.Sighting{}

	// Get row from DB, then copy data into sighting struct
	err := m.DB.QueryRow(stmt, id).Scan(&s.Index,
		&s.UserID,
		&s.Datetime,
		&s.Season,
		&s.City,
		&s.State,
		&s.Country,
		&s.Shape,
		&s.Duration,
		&s.Latitude,
		&s.Longitude,
		&s.Sighted)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *SightingModel) GetByState(state string) ([]*models.Sighting, error) {
	stmt := `SELECT * FROM sightings WHERE state = $1`

	rows, err := m.DB.Query(stmt, state)
	if err != nil {
		return nil, err
	}

	// Make sure to close up this resultset
	defer rows.Close()

	sightings := []*models.Sighting{}

	for rows.Next() {
		s := &models.Sighting{}

		err := rows.Scan(&s.Index,
			&s.UserID,
			&s.Datetime,
			&s.Season,
			&s.City,
			&s.State,
			&s.Country,
			&s.Shape,
			&s.Duration,
			&s.Latitude,
			&s.Longitude)

		if err != nil {
			return nil, err
		}

		sightings = append(sightings, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sightings, nil
}
