package app

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID uuid.UUID `gorm:"type:char(36);primary_key"`
}

func (author *Author) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}

func (author *Article) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.NewV4())
}

type Author struct {
	Base

	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Articles  []Article `json:"articles"`
}

type Article struct {
	Base

	Title    string `json:"title"`
	Body     string `json:"body"`
	Author   string `json:"author"`
	Rate     int    `json:"rate"`
	CreateAt int64  `json:"created_at"`
}
