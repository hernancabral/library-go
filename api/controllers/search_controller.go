package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hernancabral/Library/api/models"
	"net/http"
	"strconv"
	"strings"
)
// todo: when the result is epty return 404?
func (server *Server) GetBooksByPublisher(c *gin.Context) {

	// clear previous error if any
	errList = map[string]string{}

	publisher := c.Param("publisher")
	if publisher == "" {
		errList["invalid_publisher"] = "invalid publisher"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}

	books, err := models.FindBooksByPublisher(server.DB, publisher)

	if err != nil {
		errList["No_books"] = "No books Found"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": books,
	})
}

func (server *Server) GetBooksByYear(c *gin.Context) {

	// clear previous error if any
	errList = map[string]string{}

	years := c.Param("years")
	yearsSplitted := strings.Split(years, "-")
	yearFrom, err1 := strconv.Atoi(yearsSplitted[0])
	yearTo, err2 := strconv.Atoi(yearsSplitted[1])
	if years == "" || err1 != nil || err2 != nil {
		errList["invalid_years"] = "invalid years"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}

	books, err := models.FindBooksByYear(server.DB, yearFrom, yearTo)

	if err != nil {
		errList["No_books"] = "No books Found"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": books,
	})
}

func (server *Server) GetBooksByPublisherAndYear(c *gin.Context) {

	// clear previous error if any
	errList = map[string]string{}

	publisher := c.Param("publisher")
	years := c.Param("years")
	yearsSplitted := strings.Split(years, "-")
	yearFrom, err1 := strconv.Atoi(yearsSplitted[0])
	yearTo, err2 := strconv.Atoi(yearsSplitted[1])

	if publisher == "" || years == "" || err1 != nil || err2 != nil {
		errList["invalid_years"] = "invalid years"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}

	books, err := models.FindBooksByPublisherAndYear(server.DB, publisher, yearFrom, yearTo)

	if err != nil {
		errList["No_books"] = "No books Found"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": books,
	})
}

func (server *Server) GetBooksByKeyword(c *gin.Context) {

	// clear previous error if any
	errList = map[string]string{}

	keyword := c.Param("keyword")
	if keyword == "" {
		errList["invalid_publisher"] = "invalid publisher"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}

	books, err := models.FindBooksByKeyword(server.DB, keyword)

	if err != nil {
		errList["No_books"] = "No books Found"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": books,
	})
}

func (server *Server) GetBooksByTitle(c *gin.Context) {

	// clear previous error if any
	errList = map[string]string{}

	title := c.Param("title")
	if title == "" {
		errList["invalid_title"] = "invalid title"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}

	books, err := models.FindBooksByTitle(server.DB, title)

	if err != nil {
		errList["No_books"] = "No books Found"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": books,
	})
}

func (server *Server) GetBooksByAuthor(c *gin.Context) {

	// clear previous error if any
	errList = map[string]string{}

	author := c.Param("author")
	if author == "" {
		errList["invalid_publisher"] = "invalid publisher"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}

	books, err := models.FindBooksByAuthor(server.DB, author)

	if err != nil {
		errList["No_books"] = "No books Found"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": books,
	})
}
