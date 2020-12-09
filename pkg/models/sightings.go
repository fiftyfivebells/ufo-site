package postgresql

import (
	"database/sql"

	"stephenbell.dev/ufo-site/pkg/models"
)


type SightingsModel struct {
	DB *sql.DB
}

func (m *SightingsModel) Insert(userID int, datetime time.Time, season, city, state, country, shape string, duration int, latitude, longitude float64) (int, error) {
	return 0, nil
}	

func (m *SightingsModel) Get(id int) (*models.Sighting, error) {
	return nil, nil
}


