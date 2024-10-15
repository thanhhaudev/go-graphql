package model

import (
	"gorm.io/gorm"
	"time"
)

type Author struct {
	ID        int     `gorm:"primaryKey"`
	Name      string  `json:"name"`
	Books     []*Book `gorm:"many2many:author_books" json:"books"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (a Author) TableName() string {
	return "authors"
}

type AuthorInput struct {
	Name string `json:"name"`
}
