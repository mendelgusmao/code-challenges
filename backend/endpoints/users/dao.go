package users

import (
	"database/sql"
)

const (
	sqlFields      = "id, email, full_name, telephone, password"
	sqlFindByID    = "SELECT " + sqlFields + " FROM users WHERE id = ?"
	sqlFindByEmail = "SELECT " + sqlFields + " FROM users WHERE email = ?"
)

type dao struct {
	db *sql.DB
}

func newDAO(db *sql.DB) *dao {
	return &dao{db: db}
}

func (d *dao) findByID(id int64) (*User, error) {
	row := d.db.QueryRow(sqlFindByID, id)
	user, err := scanOne(row)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (d *dao) findByEmail(email string) (*User, error) {
	row := d.db.QueryRow(sqlFindByEmail, email)
	user, err := scanOne(row)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (d *dao) create(u *User) error {
	query := "INSERT INTO users (email, full_name, telephone, password) VALUES (?, ?, ?, ?)"

	result, err := d.db.Exec(query,
		u.Email,
		u.FullName,
		u.Telephone,
		u.Password,
	)

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

func (d *dao) update(u *User) error {
	query := "UPDATE users SET email = ?, full_name = ?, telephone = ?, password = ? WHERE id = ?"

	_, err := d.db.Exec(query,
		u.Email,
		u.FullName,
		u.Telephone,
		u.Password,
		u.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func scanOne(row *sql.Row) (*User, error) {
	var user User

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.FullName,
		&user.Telephone,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
