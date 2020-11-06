package models

import (
	"github.com/jinzhu/gorm"
)

type GenresResult struct {
	Genre    string
	Quantity int
}

func FindAllGenres(db *gorm.DB) (*[]GenresResult, error) {
	var err error
	genres := []GenresResult{}
	err = db.Table("books").Select("genre, count(genre) as quantity").Group("genre").Scan(&genres).Error
	if err != nil {
		return &[]GenresResult{}, err
	}
	return &genres, err
}
