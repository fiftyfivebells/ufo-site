package postgresql

import (
	"database/sql"

	"stephenbell.dev/ufo-site/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, pass string) {
	return
}

func (m *UserModel) Authenticate(email, pass string) (int, error) {
	return 0, nil
}

func (m *UserModel) Get(email, pass string) (*models.User, error) {
	return nil, nil
}
