package model

type BorrowerInput struct {
	ID int `json:"id"`
}

type Borrower struct {
	ID    int     `gorm:"primaryKey" json:"id"`
	Name  string  `json:"name"`
	Books []*Book `gorm:"foreignKey:BorrowerID" json:"books"`
}
