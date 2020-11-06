package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type Book struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:255;not null;unique" json:"title"`
	Author1   string    `gorm:"size:255;not null" json:"author1"`
	Author2   string    `gorm:"size:255;null;" json:"author2"`
	Author3   string    `gorm:"size:255;null;" json:"author3"`
	Pages     uint      `gorm:"size:255;not null;" json:"pages"`
	ISBN      string    `gorm:"size:30;not null;unique" json:"isbn"`
	Year      int       `gorm:"not null;" json:"year"`
	Genre     string    `gorm:"size:100;not null" json:"genre"`
	Publisher string    `gorm:"size:255;not null" json:"publisher"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at, omitempty"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at, omitempty"`
}

type BookRequest struct {
	Title     string `gorm:"size:255;not null;unique" json:"title"`
	Author1   string `gorm:"size:255;not null" json:"author1"`
	Author2   string `gorm:"size:255;null;" json:"author2"`
	Author3   string `gorm:"size:255;null;" json:"author3"`
	Pages     uint   `gorm:"size:255;not null;" json:"pages"`
	ISBN      string `gorm:"size:30;not null;unique" json:"isbn"`
	Year      int    `gorm:"not null;" json:"year"`
	Genre     string `gorm:"size:100;not null" json:"genre"`
	Publisher string `gorm:"size:255;not null" json:"publisher"`
}

func Validate(book *BookRequest) map[string]string {
	var errorMessages = make(map[string]string)
	var err error
	if book.Title == "" {
		err = errors.New("required title")
		errorMessages["required_title"] = err.Error()
	}
	if book.Author1 == "" {
		err = errors.New("required author")
		errorMessages["required_author"] = err.Error()
	}
	if book.Pages < 1 {
		err = errors.New("invalid page number")
		errorMessages["invalid_page_number"] = err.Error()
	}
	if book.ISBN == "" {
		err = errors.New("required isbn")
		errorMessages["required_isbn"] = err.Error()
	}
	if book.Genre == "" {
		err = errors.New("required genre")
		errorMessages["required_genre"] = err.Error()
	}
	if book.Publisher == "" {
		err = errors.New("required publisher")
		errorMessages["required_publisher"] = err.Error()
	}

	return errorMessages
}

func SaveBook(db *gorm.DB, book *Book) (*Book, error) {

	var err error
	err = db.Debug().Create(&book).Error
	if err != nil {
		return &Book{}, err
	}
	return book, nil
}

func FindAllBooks(db *gorm.DB) (*[]Book, error) {
	var err error
	books := []Book{}
	err = db.Debug().Model(&Book{}).Limit(100).Find(&books).Error
	if err != nil {
		return &[]Book{}, err
	}
	return &books, err
}

func FindBookByID(db *gorm.DB, uid uint32) (*Book, error) {
	var err error
	var book Book
	err = db.Debug().First(&book, uid).Error
	if err != nil {
		return &Book{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Book{}, errors.New("book not found")
	}
	return &book, err
}

func UpdateBook(db *gorm.DB, book *Book, uid uint32) (*Book, error) {
	err := db.Debug().Save(book).Error

	if err != nil {
		log.Fatal(errors.New("error updating book"))
	}

	// Check if it was saved correctly
	var b Book
	err = db.Debug().First(&b, uid).Error
	if db.Error != nil {
		return nil, db.Error
	}
	return &b, nil
}

func DeleteBook(db *gorm.DB, uid uint32) (bool, error) {

	db = db.Debug().Delete(&Book{}, uid)

	if db.Error != nil {
		return false, db.Error
	}

	return true, nil
}
