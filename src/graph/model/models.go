package model

type Book struct {
	ID      uint     `gorm:"primaryKey" json:"id"`
	Title   string   `json:"title"`
	Authors []Author `json:"authors"`
}

type Author struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type Borrower struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Name          string `json:"name"`
	BorrowedBooks []Book `gorm:"foreignKey:BorrowerID" json:"borrowedBooks"`
}
