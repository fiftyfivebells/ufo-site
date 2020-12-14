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
	return 0, nil
}

func (m *UserModel) Get(email, pass string) (*models.User, error) {
	return nil, nil
}
