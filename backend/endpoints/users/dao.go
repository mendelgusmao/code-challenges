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

func (d *dao) findByID(id int64) (*User, error) {
	query := "SELECT id, email FROM users WHERE id = ?"

	row := d.db.QueryRow(query, id)

	user, err := scanOne(row)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (d *dao) create(u *User) error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"

	result, err := d.db.Exec(query, u.Email, u.Password)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	u.ID = id

	return nil
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
