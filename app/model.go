package app

import (
	"golang.org/x/crypto/bcrypt"
)

// ID       string    `json:"id" gorm:"primarykey"`
type Author struct {
	Id        string    `json:"id" gorm:"primarykey"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Articles  []Article `json:"articles"`
}

type Article struct {
	Id       string `json:"id" gorm:"primarykey"`
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

func (a Author) CheckPasswordHarsh(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
	return err == nil
}
