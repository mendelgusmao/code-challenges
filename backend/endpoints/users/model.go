package users

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 10

type User struct {
	ID                           int64      `json:"id"`
	Email                        string     `json:"email"`
	FullName                     string     `json:"full_name"`
	Telephone                    string     `json:"telephone"`
	Password                     string     `json:"password,omitempty"`
	PasswordResetToken           *string    `json:"-"`
	PasswordResetTokenExpiration *time.Time `json:"-"`
}

type UserRequest User

func (u *User) apply(r *UserRequest) {
	u.Email = r.Email
	u.FullName = r.FullName
	u.Telephone = r.Telephone
	u.PasswordResetToken = r.PasswordResetToken
	u.PasswordResetTokenExpiration = r.PasswordResetTokenExpiration

	if r.Password != "" {
		u.Password = r.Password
	}
}

func (u *User) authenticate(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func (u *User) validPasswordResetToken() bool {
	if u.PasswordResetTokenExpiration == nil {
		return false
	}

	expiration := *u.PasswordResetTokenExpiration

	return !expiration.IsZero() && time.Now().UTC().Before(expiration)
}

func (u *User) filtered() User {
	return User{
		ID:        u.ID,
		Email:     u.Email,
		FullName:  u.FullName,
		Telephone: u.Telephone,
	}
}

func (u *UserRequest) validate() []error {
	return nil
}

func (u *UserRequest) encryptPassword() error {
	if u.Password == "" {
		return nil
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcryptCost)

	if err != nil {
		return err
	}

	u.Password = string(hash)

	return nil
}
