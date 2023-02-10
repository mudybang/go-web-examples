package model

import (
	"github.com/mudybang/go-web-examples/goginapi/database"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title      string `gorm:"size:255;not null;unique" json:"title"`
	Content    string `gorm:"size:255;not null;" json:"content"`
	UserID     uint
	Categories []Category `gorm:"many2many:post_categories;"`
}

func (post *Post) Save() (*Post, error) {
	err := database.Database.Create(&post).Error
	if err != nil {
		return &Post{}, err
	}
	return post, nil
}
