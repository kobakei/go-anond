package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Comment struct {
	gorm.Model
	ArticleId uint `gorm:"index"`
	Body      string
}
