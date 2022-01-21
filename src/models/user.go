package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
	IsAdmin  bool   `json:"-"`
}

func (u *User) SetPassword(password string) {
	makedHashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	u.Password = makedHashPassword
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(u.Password, []byte(password))
}
