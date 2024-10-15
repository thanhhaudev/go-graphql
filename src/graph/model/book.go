package model

import (
	"gorm.io/gorm"
	"time"
)

type BookInput struct {
	Title     string   `json:"title"`
	AuthorIds []string `json:"authorIds"`
}

type Book struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	Rating      int       `json:"rating"`
	Authors     []*Author `gorm:"many2many:author_books" json:"authors"`
	PublishAt   time.Time `json:"publishAt"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (b Book) TableName() string {
	return "books"
}
