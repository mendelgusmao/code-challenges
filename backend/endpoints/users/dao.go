package users

import (
	"database/sql"
	"time"

	"bitbucket.org/mendelgusmao/me_gu/backend/config"
)

const (
	sqlFields      = "id, email, full_name, telephone, password, password_reset_token, password_reset_token_expiration"
	sqlFindByID    = "SELECT " + sqlFields + " FROM users WHERE id = ?"
	sqlFindByEmail = "SELECT " + sqlFields + " FROM users WHERE email = ?"
	sqlFindByToken = "SELECT " + sqlFields + " FROM users WHERE password_reset_token = ?"
	sqlInsert      = "INSERT INTO users (email, full_name, telephone, password) VALUES (?, ?, ?, ?)"
	sqlUpdate      = "UPDATE users SET email = ?, full_name = ?, telephone = ?, password = ?, " +
		"password_reset_token = ?, password_reset_token_expiration = ? WHERE id = ?"
	sqlUpdateToken = "UPDATE users SET password_reset_token = ?, password_reset_token_expiration = ? WHERE email = ?"
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

func (d *dao) findByToken(token string) (*User, error) {
	row := d.db.QueryRow(sqlFindByToken, token)
	user, err := scanOne(row)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (d *dao) create(u *User) error {
	result, err := d.db.Exec(sqlInsert,
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
	_, err := d.db.Exec(sqlUpdate,
		u.Email,
		u.FullName,
		u.Telephone,
		u.Password,
		u.PasswordResetToken,
		u.PasswordResetTokenExpiration,
		u.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (d *dao) updatePasswordResetToken(email, token string) error {
	expiration := time.Now().UTC().Add(config.Backend.PasswordResetExpiration)

	_, err := d.db.Exec(sqlUpdateToken,
		token,
		expiration,
		email,
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
		&user.PasswordResetToken,
		&user.PasswordResetTokenExpiration,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
