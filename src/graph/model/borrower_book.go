package model

import (
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const (
	BorrowedStatus = "borrowed"
	ReturnedStatus = "returned"
)

type BorrowBookInput struct {
	BorrowerID int       `json:"borrowerId"`
	BookID     int       `json:"bookId"`
	Quantity   int       `json:"quantity"`
	DueDate    time.Time `json:"dueDate"`
}

func (b *BorrowBookInput) Validate() error {
	if b.Quantity <= 0 {
		return fmt.Errorf("quantity must be greater than 0")
	}

	if b.DueDate.IsZero() {
		return fmt.Errorf("due date must be provided")
	}

	if b.DueDate.Before(time.Now()) {
		return fmt.Errorf("due date must be in the future")
	}

	return nil
}

type BorrowerBook struct {
	ID         int            `gorm:"primaryKey" json:"id"`
	BorrowerID int            `gorm:"not null" json:"borrowerId"`
	Borrower   *Borrower      `gorm:"foreignKey:BorrowerID;references:ID" json:"borrower"`
	BookID     int            `gorm:"not null" json:"bookId"`
	Book       *Book          `gorm:"foreignKey:BookID;references:ID" json:"book"`
	BorrowedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"borrowedAt"`
	ReturnedAt *time.Time     `json:"returnedAt"`
	DueDate    datatypes.Date `gorm:"type:DATE;not null;" json:"dueDate"`
	Status     string         `gorm:"type:VARCHAR(10);not null;check:status IN ('borrowed', 'returned');default:'borrowed'" json:"status"`
	Quantity   int            `gorm:"not null;default:1" json:"quantity"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (bb BorrowerBook) TableName() string {
	return "borrower_books"
}
