package users

import "golang.org/x/crypto/bcrypt"

const bcryptCost = 10

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
}

func (u *User) validate() []error {
	return nil
}

func (u *User) encryptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcryptCost)

	if err != nil {
		return err
	}

	u.Password = string(hash)

	return nil
}
