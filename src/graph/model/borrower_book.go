package model

import (
	"gorm.io/gorm"
	"time"
)

type BorrowersBooks struct {
	BorrowerID int       `gorm:"primaryKey" json:"borrowerId"`
	BookID     int       `gorm:"primaryKey" json:"bookId"`
	BorrowedAt time.Time `json:"borrowedAt"`
	ReturnedAt time.Time `json:"returnedAt"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (bb BorrowersBooks) TableName() string {
	return "borrowers_books"
}
