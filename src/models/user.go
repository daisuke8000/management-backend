package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Point    int    `json:"point"`
	Rank     int    `json:"rank"`
	Password []byte `json:"-"`
	IsAdmin  bool   `json:"-"`
}

func (u *User) SetPassword(password string) {
	makedHashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	u.Password = makedHashPassword
}
