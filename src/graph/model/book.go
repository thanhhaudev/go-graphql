package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type CreateBookInput struct {
	Title     string `json:"title"`
	AuthorIds []int  `json:"authorIds"`
	PublishAt time.Time
	Quantity  int
	Rating    float64
}

func (c *CreateBookInput) Validate() error {
	if c.Title == "" {
		return fmt.Errorf("title can't be blank")
	}

	if c.Quantity < 0 {
		return fmt.Errorf("quantity can't be negative")
	}

	if c.Rating < 0 || c.Rating > 5 {
		return fmt.Errorf("rating must be between 0 and 5")
	}

	if len(c.AuthorIds) == 0 {
		return fmt.Errorf("authors can't be blank")
	}

	if c.PublishAt.IsZero() {
		return fmt.Errorf("publishAt can't be blank")
	}

	return nil
}

func (c *CreateBookInput) ToBook() *Book {
	return &Book{
		Title:     c.Title,
		PublishAt: c.PublishAt,
		Quantity:  c.Quantity,
		Rating:    c.Rating,
	}
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

func (b *Book) SetAuthors(authors []*Author) {
	b.Authors = authors
}
