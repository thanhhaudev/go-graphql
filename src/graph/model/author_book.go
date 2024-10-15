package model

import (
	"gorm.io/gorm"
	"time"
)

type AuthorBook struct {
	AuthorID  int `gorm:"primaryKey"`
	BookID    int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (AuthorBook) TableName() string {
	return "author_books"
}
