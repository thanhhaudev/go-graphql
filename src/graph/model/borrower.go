package model

type Borrower struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Books []Book `gorm:"foreignKey:BorrowerID" json:"books"`
}
