package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hernancabral/Library/api/models"
	"net/http"
)

func (server *Server) GetGenres(c *gin.Context) {

	// clear previous error if any
	errList = map[string]string{}

	genres, err := models.FindAllGenres(server.DB)
	if err != nil {
		errList["No_genres"] = "No genres Found"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": genres,
	})
}
