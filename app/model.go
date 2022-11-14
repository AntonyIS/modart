package app

import "github.com/jinzhu/gorm"

type Author struct {
	gorm.Model
	Id        string    `json:"id"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"eamil"`
	Password  string    `json:"password"`
	Articles  []Article `json:"article"`
}

type Article struct {
	gorm.Model
	Id       string `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Author   string `json:"author"`
	Rate     int    `json:"rate"`
	CreateAt int64  `json:"created_at"`
}
