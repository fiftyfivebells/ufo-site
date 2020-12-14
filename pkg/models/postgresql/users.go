package postgresql

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"stephenbell.dev/ufo-site/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, pass string) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (username, email, hashed_pass, created)
VALUES($1, $2, $3, NOW())`

	_, err = m.DB.Exec(stmt, name, email, string(hashedPass))
	if err != nil {
		var pgSQLError *pq.Error
		if errors.As(err, &pgSQLError) {
			if strings.Contains(pgSQLError.Message, "users_uc_email") {
				return models.ErrDuplicateEmail
			}
		}
		return err
	}

	return nil
}

func (m *UserModel) Authenticate(email, pass string) (int, error) {
	var id int
	var hashedPass []byte

	stmt := "SELECT id, hashed_pass FROM users WHERE email = $1 AND active = TRUE"
	row := m.DB.QueryRow(stmt, email)

	err := row.Scan(&id, &hashedPass)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, models.ErrInvalidCredentials
		} else {
			return -1, err
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPass, []byte(pass))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return -1, models.ErrInvalidCredentials
		} else {
			return -1, err
		}
	}

	return id, nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
	u := &models.User{}

	stmt := "SELECT id, username, email, created, active FROM users WHERE id = $1"
	err := m.DB.QueryRow(stmt, id).Scan(&u.ID, u.Username, u.Email, u.Created, u.Active)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return u, nil
}
