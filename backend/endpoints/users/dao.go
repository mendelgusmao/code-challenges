package users

import (
	"database/sql"
)

type dao struct {
	db *sql.DB
}

func newDAO(db *sql.DB) *dao {
	return &dao{db: db}
}

func (d *dao) findByID(id int) (*User, error) {
	query := "SELECT id, email FROM users WHERE id = ?"

	row := d.db.QueryRow(query, id)

	user, err := scanOne(row)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func scanOne(row *sql.Row) (*User, error) {
	var user User

	err := row.Scan(
		&user.ID,
		&user.Email,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
