package model

// AuthorBook represents the many-to-many relationship between authors and books.
// this is the join table, so no need to define any fields other than the foreign keys.
type AuthorBook struct {
	AuthorID int `gorm:"primaryKey"`
	BookID   int `gorm:"primaryKey"`
}

func (AuthorBook) TableName() string {
	return "authors_books"
}
