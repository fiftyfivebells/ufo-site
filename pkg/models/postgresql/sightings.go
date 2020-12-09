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
	stmt := `INSERT INTO sightings (user_id, datetime, season, city, state, country, shape, duration, latitude, longitude) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := m.DB.Exec(stmt, userID, datetime, season, city, state, country, shape, duration, lat, long)

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return -1, err
	}

	return int(id), nil
}

func (m *SightingModel) Get(id int) (*models.Sighting, error) {
	return nil, nil
}
