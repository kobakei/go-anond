package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Article struct {
	gorm.Model
	Title string `gorm:"size:255"`
	Body  string

	Comments []Comment
}

func (a *Article) Summary() string {
	runes := []rune(a.Body)
	return string(runes[0:10])
}
