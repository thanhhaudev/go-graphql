package model

import (
	"gorm.io/gorm"
	"time"
)

type CreateBookInput struct {
	Title     string   `json:"title"`
	AuthorIds []string `json:"authorIds"`
	PublishAt time.Time
	Quantity  int
	Rating    float64
}

type UpdateBookInput struct {
	Title     *string  `json:"title"`
	AuthorIds []string `json:"authorIds"`
	PublishAt *time.Time
	Quantity  *int
	Rating    *float64
}

type Book struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	Rating      float64   `json:"rating"`
	Authors     []*Author `gorm:"many2many:authors_books" json:"authors"`
	PublishAt   time.Time `json:"publishAt"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (b Book) TableName() string {
	return "books"
}
