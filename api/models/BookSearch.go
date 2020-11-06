package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

func FindBooksByPublisher(db *gorm.DB, publisher string) (*[]Book, error) {
	var err error
	var books []Book
	err = db.Debug().Where("UPPER(publisher) LIKE UPPER(?)", "%" + publisher + "%").Find(&books).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("book not found")
	}
	return &books, err
}

func FindBooksByAuthor(db *gorm.DB, author string) (*[]Book, error) {
	var err error
	var books []Book
	err = db.Debug().Where("UPPER(author1) LIKE UPPER(?)", "%" +author+ "%").Or("UPPER(author2) LIKE UPPER(?)", "%" +author+ "%").Or("UPPER(author3) LIKE UPPER(?)", "%" +author+ "%").Find(&books).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err){
		return nil, errors.New("book not found")
	}
	return &books, err
}

func FindBooksByKeyword(db *gorm.DB, keyword string) (*[]Book, error) {
	var err error
	var books []Book
	err = db.Debug().Where("UPPER(title) LIKE UPPER(?)", "%" + keyword + "%").Or("UPPER(author1) LIKE UPPER(?)", "%" + keyword + "%").Or("UPPER(author2) LIKE UPPER(?)", "%" + keyword + "%").Or("UPPER(author3) LIKE UPPER(?)", "%" + keyword + "%").Or("UPPER(genre) LIKE UPPER(?)", "%" + keyword + "%").Or("UPPER(publisher) LIKE UPPER(?)", "%" + keyword + "%").Find(&books).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("book not found")
	}
	return &books, err
}

func FindBooksByYear(db *gorm.DB, yearFrom, yearTo int) (*[]Book, error) {
	var err error
	var books []Book
	err = db.Debug().Find(&books, "year BETWEEN ? AND ?", yearFrom, yearTo).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("book not found")
	}
	return &books, err
}

// TODO: keyword, editorial, year, author, publisher, publisher and year -- genres author
func FindBooksByPublisherAndYear(db *gorm.DB, publisher string, yearFrom, yearTo int) (*[]Book, error) {
	var err error
	var books []Book
	err = db.Debug().Find(&books, "UPPER(publisher) LIKE UPPER(?) AND year BETWEEN ? AND ?", "%" + publisher + "%", yearFrom, yearTo).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("book not found")
	}
	return &books, err
}
