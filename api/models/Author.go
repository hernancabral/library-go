package models

import (
	"github.com/jinzhu/gorm"
)

type AuthorsResult struct {
	Author  string
	Quantity int
}

func FindAllAuthors(db *gorm.DB) (*[]AuthorsResult, error) {
	var err error
	authors := []AuthorsResult{}
	err = db.Table("books").Select("author1 as author, count(author1) as quantity").Group("author1").Scan(&authors).Error
	if err != nil {
		return &[]AuthorsResult{}, err
	}
	return &authors, err
}
