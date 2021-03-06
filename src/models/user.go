package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Model
	Name     string `json:"name" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
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
