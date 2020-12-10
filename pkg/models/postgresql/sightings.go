package postgresql

import (
	"database/sql"
	"time"

	"stephenbell.dev/ufo-site/pkg/models"
)

type SightingModel struct {
	DB *sql.DB
}

func (m *SightingModel) Insert(userID int, datetime time.Time, season, city, state, country, shape string, duration int, lat, long float64) (int, error) {
	// stmt := `INSERT INTO sightings  VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	query := `INSERT INTO sightings (user_id, datetime, season, city, state, country, shape, duration, latitude, longitude) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING index`

	stmt, err := m.DB.Prepare(query)

	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRow(userID, datetime, season, city, state, country, shape, duration, lat, long).Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (m *SightingModel) Get(id int) (*models.Sighting, error) {
	return nil, nil
}
