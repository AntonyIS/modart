package app

import (
	"golang.org/x/crypto/bcrypt"
)

type Author struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Articles  []Article `json:"articles"`
}

type Article struct {
	ID       string `json:"id"`
	AuthorID string `json:"author_id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Author   string `json:"author"`
	Rate     int    `json:"rate"`
	CreateAt int64  `json:"created_at"`
}

func (a Author) GenerateHashPassord() (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
