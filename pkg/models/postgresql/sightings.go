package postgresql

import (
	"database/sql"
	"time"

	"stephenbell.dev/ufo-site/pkg/models"
)

type SightingModel struct {
	DB *sql.DB
}

func (m *SightingModel) Insert(userID int, datetime time.Time, season, city, state, country, shape string, duration int, latitude, longitude float64) (int, error) {
	return 0, nil
}

func (m *SightingModel) Get(id int) (*models.Sighting, error) {
	return nil, nil
}
