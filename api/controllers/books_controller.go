package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hernancabral/Library/api/models"
	"github.com/hernancabral/Library/api/utils"
	"github.com/hernancabral/Library/api/utils/formaterror"
	"net/http"
	"strconv"
	"time"
)

func (server *Server) PostBook(c *gin.Context) {

	// clear previous error if any
	errList = map[string]string{}

	var input models.BookRequest

	// If there was an error parsing json body to Book body struct, request is aborted with Bad Request status code.
	if err := c.BindJSON(&input); err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}

	// Check if the body is valid
	if len(models.Validate(&input)) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  models.Validate(&input),
		})
		return
	}

	newBook := models.Book{
		Title:     input.Title,
		Author1:   input.Author1,
		Author2:   input.Author2,
		Author3:   input.Author3,
		Pages:     input.Pages,
		ISBN:      input.ISBN,
		Year:      input.Year,
		Genre:     input.Genre,
		Publisher: input.Publisher,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	updatedBook, err := models.SaveBook(server.DB, &newBook)
	if err != nil {
		errList := formaterror.FormatError(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusCreated,
		"response": updatedBook,
	})
}

func (server *Server) GetBooks(c *gin.Context) {

	// clear previous error if any
	errList = map[string]string{}

	books, err := models.FindAllBooks(server.DB)
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

func (server *Server) GetBookById(c *gin.Context) {

	// clear previous error if any
	errList = map[string]string{}

	bookID := c.Param("id")

	uid, err := strconv.ParseUint(bookID, 10, 32)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}

	bookFound, err := models.FindBookByID(server.DB, uint32(uid))
	if err != nil {
		errList["No_book"] = "No Book Found"
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": bookFound,
	})
}

func (server *Server) UpdateBook(c *gin.Context) {

	// clear previous error if any
	errList = map[string]string{}

	var input models.BookRequest

	// If there was an error parsing json body to Book body struct, request is aborted with Bad Request status code.
	if err := c.BindJSON(&input); err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}

	// Check if the id is valid
	bookID := c.Param("id")
	id, err := strconv.ParseUint(bookID, 10, 32)

	if bookID == "" || !utils.IsNumeric(bookID) || err != nil {
		errList["invalid_id"] = "invalid_id"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}

	// Check if the body is valid
	if len(models.Validate(&input)) > 0 {
		errList["invalid_id"] = "invalid id"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}

	// Check if the books exists
	_, err = models.FindBookByID(server.DB, uint32(id))
	if err != nil {
		errList["book_not_exist"] = "the book does not exist"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.NotFound,
			"error":  errList,
		})
		return
	}

	newBook := models.Book{
		ID:        uint32(id),
		Title:     input.Title,
		Author1:   input.Author1,
		Author2:   input.Author2,
		Author3:   input.Author3,
		Pages:     input.Pages,
		ISBN:      input.ISBN,
		Year:      input.Year,
		Genre:     input.Genre,
		Publisher: input.Publisher,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	updatedBook, err := models.UpdateBook(server.DB, &newBook, uint32(id))
	if err != nil {
		errList := formaterror.FormatError(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": updatedBook,
	})
}

func (server *Server) DeleteBook(c *gin.Context) {

	// clear previous error if any
	errList = map[string]string{}

	bookID := c.Param("id")

	uid, err := strconv.ParseUint(bookID, 10, 32)
	if err != nil {
		errList["Invalid_request"] = "Invalid Request"
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  errList,
		})
		return
	}

	book, err := models.DeleteBook(server.DB, uint32(uid))
	if err != nil {
		errList["error_deleting"] = "Error deleting book"
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  errList,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": book,
	})
}
