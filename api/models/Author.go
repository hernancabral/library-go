package models

import (
	"github.com/jinzhu/gorm"
)

type AuthorsResult struct {
	Author   string
	Quantity int
}

func FindAllAuthors(db *gorm.DB) (*[]AuthorsResult, error) {
	var err error
	authors1 := []AuthorsResult{}
	authors2 := []AuthorsResult{}
	authors3 := []AuthorsResult{}

	err = db.Table("books").Select("author1 as author, count(author1) as quantity").Group("author1").Scan(&authors1).Error
	if err != nil {
		return &[]AuthorsResult{}, err
	}
	err = db.Table("books").Select("author2 as author, count(author2) as quantity").Group("author2").Scan(&authors2).Error
	if err != nil {
		return &[]AuthorsResult{}, err
	}
	err = db.Table("books").Select("author3 as author, count(author3) as quantity").Group("author3").Scan(&authors3).Error
	if err != nil {
		return &[]AuthorsResult{}, err
	}

	allAuthors := append(authors1, authors2...)
	allAuthors = append(allAuthors, authors3...)

	authorMap := make(map[string]int)
	for _, v := range allAuthors {
		if v.Author != "" {
			authorMap[v.Author] += v.Quantity
		}
	}

	result := []AuthorsResult{}
	for a, q := range authorMap {
		result = append(result, AuthorsResult{Author: a, Quantity: q})
	}
	return &result, err
}
