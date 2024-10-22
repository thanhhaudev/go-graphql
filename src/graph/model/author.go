package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        int     `gorm:"primaryKey"`
	Name      string  `json:"name"`
	Books     []*Book `gorm:"many2many:authors_books" json:"books"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (a Author) TableName() string {
	return "authors"
}

type CreateAuthorInput struct {
	Name  string `json:"name"`
	Books []int  `json:"books"`
}

func (c *CreateAuthorInput) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("name can't be blank")
	}

	return nil
}

func (c *CreateAuthorInput) ToAuthor() *Author {
	return &Author{
		Name: c.Name,
	}
}

type UpdateAuthorInput struct {
	Name *string `json:"name"`
}
