package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hernancabral/Library/api/models"
	"net/http"
)

func (server *Server) GetAuthors(c *gin.Context) {

	// clear previous error if any
	errList = map[string]string{}

	authors, err := models.FindAllAuthors(server.DB)
	if err != nil {
		errList["no_authors"] = "no authors found"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": authors,
	})
}
