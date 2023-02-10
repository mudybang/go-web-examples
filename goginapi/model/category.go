package model

import (
	"github.com/mudybang/go-web-examples/goginapi/database"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"size:255;not null;" json:"name"`
	Slug string `gorm:"size:255;not null;unique" json:"slug"`
}

func (raw *Category) Save() (*Category, error) {
	err := database.Database.Create(&raw).Error
	if err != nil {
		return &Category{}, err
	}
	return raw, nil
}
