package users

import "golang.org/x/crypto/bcrypt"

const bcryptCost = 10

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	Telephone string `json:"telephone"`
	Password  string `json:"password,omitempty"`
}

type UserRequest User

func (u *User) apply(r *UserRequest) {
	u.Email = r.Email
	u.FullName = r.FullName
	u.Telephone = r.Telephone

	if r.Password != "" {
		u.Password = r.Password
	}
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
