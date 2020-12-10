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

func (m *SightingModel) Insert(userID int, datetime time.Time, season, city, state, country, shape string, duration int, lat, long float64) (int, error) {

	query := `INSERT INTO sightings (user_id, datetime, season, city, state, country, shape, duration, latitude, longitude) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING index`

	stmt, err := m.DB.Prepare(query)

	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	var id int
	err = stmt.QueryRow(userID,
		datetime,
		season,
		city,
		state,
		country,
		shape,
		duration,
		lat,
		long).Scan(&id)

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
		&s.Longitude)

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

	rows, err := m.DB.Query(stmt)
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
