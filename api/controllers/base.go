package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hernancabral/Library/api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

var errList = make(map[string]string)

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	var err error
	server.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error connecting to postgres:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	// database migration
	server.DB.Debug().AutoMigrate(
		&models.Book{},
	)

	server.Router = gin.Default()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
