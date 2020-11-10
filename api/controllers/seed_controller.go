package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hernancabral/Library/api/seed"
	"net/http"
)

func (server *Server) Seed(c *gin.Context) {

	// clear previous error if any
	errList = map[string]string{}

	books, err := seed.Load(server.DB)
	if err != nil {
		errList["error_seeding"] = "Error seeding DB"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": books,
	})
}
