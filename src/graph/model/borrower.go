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

type UpdateBorrowerInput struct {
	Name      *string    `json:"name"`
	Address   *string    `json:"address"`
	TelNumber *string    `json:"telNumber"`
	BirthDate *time.Time `json:"birthDate"`
}

func (u *UpdateBorrowerInput) Validate() error {
	if u.Name != nil && *u.Name == "" {
		return fmt.Errorf("name is required")
	}

	if u.Address != nil && *u.Address == "" {
		return fmt.Errorf("address is required")
	}

	if u.TelNumber != nil && *u.TelNumber == "" {
		return fmt.Errorf("telNumber is required")
	}

	if u.BirthDate != nil && u.BirthDate.IsZero() {
		return fmt.Errorf("birthDate is required")
	}

	return nil
}

func (u *UpdateBorrowerInput) Patch(b *Borrower) {
	if u.Name != nil {
		b.Name = *u.Name
	}

	if u.Address != nil {
		b.Address = *u.Address
	}

	if u.TelNumber != nil {
		b.TelNumber = *u.TelNumber
	}

	if u.BirthDate != nil {
		b.BirthDate = datatypes.Date(*u.BirthDate)
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
