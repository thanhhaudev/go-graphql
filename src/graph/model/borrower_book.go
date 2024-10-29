package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"time"
)

type BorrowerBook struct {
	ID         int            `gorm:"primaryKey" json:"id"`
	BorrowerID int            `gorm:"not null" json:"borrowerId"`
	Borrower   *Borrower      `gorm:"foreignKey:BorrowerID;references:ID" json:"borrower"`
	BookID     int            `gorm:"not null" json:"bookId"`
	Book       *Book          `gorm:"foreignKey:BookID;references:ID" json:"book"`
	BorrowedAt time.Time      `json:"borrowedAt"`
	ReturnedAt *time.Time     `json:"returnedAt"`
	DueDate    datatypes.Date `gorm:"type:DATE;not null;" json:"dueDate"`
	Status     string         `gorm:"type:VARCHAR(10);not null;check:status IN ('borrowed', 'returned');default:'borrowed'" json:"status"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (bb BorrowerBook) TableName() string {
	return "borrower_books"
}
