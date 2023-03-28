package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint
	FirstName string
	LastName  string
	Email     string
	Password  string
	IsMithril bool
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = string(hashedPassword)
}

func (user *User) CompareHashedPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
