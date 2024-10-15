package model

import (
	"gorm.io/gorm"
	"time"
)

type BorrowerInput struct {
	ID int `json:"id"`
}

type Borrower struct {
	ID        int     `gorm:"primaryKey" json:"id"`
	Name      string  `json:"name"`
	Books     []*Book `gorm:"many2many:borrowers_books" json:"books"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (b Borrower) TableName() string {
	return "borrowers"
}
