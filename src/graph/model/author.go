package model

type Author struct {
	ID    int    `gorm:"primaryKey"`
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type AuthorInput struct {
	Name string `json:"name"`
}
