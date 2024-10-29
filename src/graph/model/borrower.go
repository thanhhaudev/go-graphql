package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"fmt"
	"time"
)

type CreateBorrowerInput struct {
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	TelNumber string    `json:"telNumber"`
	BirthDate time.Time `json:"birthDate"`
}

func (c *CreateBorrowerInput) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("name is required")
	}

	if c.Address == "" {
		return fmt.Errorf("address is required")
	}

	if c.TelNumber == "" {
		return fmt.Errorf("telNumber is required")
	}

	if c.BirthDate.IsZero() {
		return fmt.Errorf("birthDate is required")
	}

	return nil
}

func (c *CreateBorrowerInput) ToBorrower() *Borrower {
	return &Borrower{
		Name:      c.Name,
		Address:   c.Address,
		TelNumber: c.TelNumber,
		BirthDate: datatypes.Date(c.BirthDate),
	}
}

type Borrower struct {
	ID        int             `gorm:"primaryKey" json:"id"`
	Name      string          `json:"name"`
	Address   string          `json:"address"`
	TelNumber string          `json:"telNumber"`
	BirthDate datatypes.Date  `gorm:"type:DATE;not null;" json:"birthDate"`
	Borrowed  []*BorrowerBook `gorm:"foreignKey:BorrowerID;references:ID" json:"borrowed"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (b Borrower) TableName() string {
	return "borrowers"
}
