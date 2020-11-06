package seed

import (
	"github.com/hernancabral/Library/api/models"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var books = []models.Book{
	{
		Title:     "The Second Sleep",
		Author1:   "Robert Harris",
		Author2:   "",
		Author3:   "",
		Pages:     448,
		ISBN:      "9781787460966",
		Year:      2020,
		Genre:     "Fiction",
		Publisher: "Cornerstone",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	},
	{
		Title:     "The Silent Patient",
		Author1:   "Alex Michaelides",
		Author2:   "",
		Author3:   "",
		Pages:     352,
		ISBN:      "9781409181637",
		Year:      2020,
		Genre:     "Fiction",
		Publisher: "Orion Publishing Co",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	},
	{
		Title:     "Abstract Art (World of Art)",
		Author1:   "Anna Moszynska",
		Author2:   "",
		Author3:   "",
		Pages:     272,
		ISBN:      "9780500204450",
		Year:      2020,
		Genre:     "Art",
		Publisher: "Thames & Hudson Ltd",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	},
	{
		Title:     "Modern Architecture (World of Art) : A Critical History",
		Author1:   "Kenneth Frampton",
		Author2:   "",
		Author3:   "",
		Pages:     736,
		ISBN:      "9780500204443",
		Year:      2020,
		Genre:     "Art",
		Publisher: "Thames & Hudson Ltd",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	},
	{
		Title:     "Normal People",
		Author1:   "Sally Rooney",
		Author2:   "",
		Author3:   "",
		Pages:     288,
		ISBN:      "9780571334650",
		Year:      2020,
		Genre:     "Romance",
		Publisher: "FABER & FABER",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	},
	{
		Title:     "Rebecca",
		Author1:   "Daphne du Maurier",
		Author2:   "",
		Author3:   "",
		Pages:     432,
		ISBN:      "9780349006574",
		Year:      2016,
		Genre:     "Romance",
		Publisher: "Virago Press Ltd",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Book{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Book{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range books {
		err = db.Debug().Model(&models.Book{}).Create(&books[i]).Error
		if err != nil {
			log.Fatalf("cannot seed books table: %v", err)
		}
	}
}