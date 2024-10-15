package seeder

import (
	"github.com/thanhhaudev/go-graphql/src/graph/model"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	Authors = []*model.Author{
		{
			Name:      "J.K. Rowling",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "George R.R. Martin",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "J.R.R. Tolkien",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Stephen King",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Agatha Christie",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Mark Twain",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Ernest Hemingway",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "F. Scott Fitzgerald",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Charles Dickens",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Jane Austen",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Leo Tolstoy",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Gabriel Garcia Marquez",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Haruki Murakami",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Toni Morrison",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Isabel Allende",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "George Orwell",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	Books = []*model.Book{
		{
			Title:       "Harry Potter and the Sorcerer's Stone",
			Description: "A young wizard embarks on his journey to defeat the dark wizard Voldemort.",
			Quantity:    100,
			Rating:      5,
			Authors:     []*model.Author{Authors[0]}, // J.K. Rowling
			PublishAt:   time.Date(1997, 6, 26, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "A Game of Thrones",
			Description: "A political struggle in the Seven Kingdoms of Westeros, involving great families and dragons.",
			Quantity:    75,
			Rating:      5,
			Authors:     []*model.Author{Authors[1]}, // George R.R. Martin
			PublishAt:   time.Date(1996, 8, 6, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "The Hobbit",
			Description: "A hobbit named Bilbo Baggins embarks on a quest to reclaim a treasure guarded by a dragon.",
			Quantity:    50,
			Rating:      5,
			Authors:     []*model.Author{Authors[2]}, // J.R.R. Tolkien
			PublishAt:   time.Date(1937, 9, 21, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "The Shining",
			Description: "A family becomes trapped in a haunted hotel during the winter.",
			Quantity:    60,
			Rating:      5,
			Authors:     []*model.Author{Authors[3]}, // Stephen King
			PublishAt:   time.Date(1977, 1, 28, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "Murder on the Orient Express",
			Description: "A detective investigates a murder on a famous train.",
			Quantity:    80,
			Rating:      5,
			Authors:     []*model.Author{Authors[4]}, // Agatha Christie
			PublishAt:   time.Date(1934, 1, 1, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "Adventures of Huckleberry Finn",
			Description: "A boy and a runaway slave journey down the Mississippi River.",
			Quantity:    40,
			Rating:      4,
			Authors:     []*model.Author{Authors[5]}, // Mark Twain
			PublishAt:   time.Date(1884, 12, 10, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "The Old Man and the Sea",
			Description: "An aging fisherman struggles to catch a giant marlin.",
			Quantity:    45,
			Rating:      5,
			Authors:     []*model.Author{Authors[6]}, // Ernest Hemingway
			PublishAt:   time.Date(1952, 9, 1, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "The Great Gatsby",
			Description: "A mysterious millionaire throws lavish parties in search of a lost love.",
			Quantity:    90,
			Rating:      5,
			Authors:     []*model.Author{Authors[7]}, // F. Scott Fitzgerald
			PublishAt:   time.Date(1925, 4, 10, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "A Tale of Two Cities",
			Description: "A story set during the French Revolution, focusing on themes of sacrifice and resurrection.",
			Quantity:    85,
			Rating:      5,
			Authors:     []*model.Author{Authors[8]}, // Charles Dickens
			PublishAt:   time.Date(1859, 4, 30, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "Pride and Prejudice",
			Description: "A romantic novel about manners, upbringing, and marriage in early 19th-century England.",
			Quantity:    70,
			Rating:      5,
			Authors:     []*model.Author{Authors[9]}, // Jane Austen
			PublishAt:   time.Date(1813, 1, 28, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "War and Peace",
			Description: "A historical novel set during the Napoleonic Wars, focusing on five aristocratic families.",
			Quantity:    30,
			Rating:      5,
			Authors:     []*model.Author{Authors[10]}, // Leo Tolstoy
			PublishAt:   time.Date(1869, 1, 1, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "One Hundred Years of Solitude",
			Description: "A multi-generational story about the Buendía family in a fictional town in Colombia.",
			Quantity:    55,
			Rating:      5,
			Authors:     []*model.Author{Authors[11]}, // Gabriel Garcia Marquez
			PublishAt:   time.Date(1967, 5, 30, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "Norwegian Wood",
			Description: "A nostalgic story of loss, sexuality, and emotional turmoil in 1960s Japan.",
			Quantity:    65,
			Rating:      5,
			Authors:     []*model.Author{Authors[12]}, // Haruki Murakami
			PublishAt:   time.Date(1987, 9, 4, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "Beloved",
			Description: "A novel about the aftereffects of slavery and the haunting of one family.",
			Quantity:    35,
			Rating:      5,
			Authors:     []*model.Author{Authors[13]}, // Toni Morrison
			PublishAt:   time.Date(1987, 9, 16, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "The House of the Spirits",
			Description: "A family saga set in Chile, blending magical realism with political events.",
			Quantity:    40,
			Rating:      5,
			Authors:     []*model.Author{Authors[14]}, // Isabel Allende
			PublishAt:   time.Date(1982, 1, 1, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Title:       "1984",
			Description: "A dystopian novel about a totalitarian regime that uses surveillance and thought control.",
			Quantity:    95,
			Rating:      5,
			Authors:     []*model.Author{Authors[15]}, // George Orwell
			PublishAt:   time.Date(1949, 6, 8, 0, 0, 0, 0, time.UTC),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	Borrowers = []*model.Borrower{
		{
			Name:      "John Smith",
			Books:     []*model.Book{Books[0], Books[1]}, // Harry Potter and A Game of Thrones
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Alice Johnson",
			Books:     []*model.Book{Books[2]}, // The Hobbit
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Michael Brown",
			Books:     []*model.Book{Books[3], Books[4]}, // The Shining and Murder on the Orient Express
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Emily Davis",
			Books:     []*model.Book{Books[5], Books[6]}, // Adventures of Huckleberry Finn and The Old Man and the Sea
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "James Wilson",
			Books:     []*model.Book{Books[7]}, // The Great Gatsby
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Sophia Martinez",
			Books:     []*model.Book{Books[8]}, // A Tale of Two Cities
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Daniel Anderson",
			Books:     []*model.Book{Books[9], Books[10]}, // Pride and Prejudice and War and Peace
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Jessica Lee",
			Books:     []*model.Book{Books[11]}, // One Hundred Years of Solitude
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Christopher Garcia",
			Books:     []*model.Book{Books[12]}, // Norwegian Wood
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Olivia Martinez",
			Books:     []*model.Book{Books[13]}, // Beloved
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "David Taylor",
			Books:     []*model.Book{Books[14]}, // The House of the Spirits
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Chloe White",
			Books:     []*model.Book{Books[15]}, // 1984
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
)

func SeedAuthors(db *gorm.DB, authors []*model.Author) {
	var count int64
	db.Model(&model.Author{}).Count(&count)

	if count == 0 {

		if err := db.Create(&authors).Error; err != nil {
			log.Fatalf("could not seed users: %v", err)
		}
		log.Println("Users table seeded.")
	} else {
		log.Println("Users table already contains data.")
	}
}

func SeedBooks(db *gorm.DB, books []*model.Book) {
	var count int64
	db.Model(&model.Book{}).Count(&count)
	if count == 0 {
		if err := db.Create(&books).Error; err != nil {
			log.Fatalf("could not seed books: %v", err)
		}
		log.Println("Books table seeded.")
	} else {
		log.Println("Books table already contains data.")
	}
}

func SeedBorrowers(db *gorm.DB, borrowers []*model.Borrower) {
	var count int64
	db.Model(&model.Borrower{}).Count(&count)
	if count == 0 {
		if err := db.Create(&borrowers).Error; err != nil {
			log.Fatalf("could not seed borrowers: %v", err)
		}
		log.Println("Borrowers table seeded.")
	} else {
		log.Println("Borrowers table already contains data.")
	}
}
