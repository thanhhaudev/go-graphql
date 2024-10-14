package model

type BookInput struct {
	Title     string   `json:"title"`
	AuthorIds []string `json:"authorIds"`
}

type Book struct {
	ID      int      `gorm:"primaryKey" json:"id"`
	Title   string   `json:"title"`
	Authors []Author `json:"authors"`
}
