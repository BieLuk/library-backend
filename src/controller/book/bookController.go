package book

import (
	"fmt"
	"github.com/BieLuk/library-backend/prometheus"
	"github.com/BieLuk/library-backend/src/apperr"
	"github.com/BieLuk/library-backend/src/dto"
	"github.com/BieLuk/library-backend/src/service/books"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

type BookController interface {
	CreateBook(c *gin.Context)
	GetBooks(c *gin.Context)
	GetBook(c *gin.Context)
}

type bookController struct {
	bookService books.BookService
}

func NewBookController(bookService books.BookService) *bookController {
	return &bookController{
		bookService: bookService,
	}
}

// CreateBook creates model.Book object in database
func (bc *bookController) CreateBook(c *gin.Context) {
	var request dto.CreateBookRequest
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			apperr.NewAppErr(apperr.BAD_REQUEST, fmt.Sprintf("cannot unmarshall request object: %v", err)))
		return
	}

	response, err := bc.bookService.CreateBook(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			apperr.NewAppErr(apperr.INTERNAL_ERROR, fmt.Sprintf("cannot create book: %v", err)))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetBooks returns list of all books in library
func (bc *bookController) GetBooks(c *gin.Context) {
	response, err := bc.bookService.GetBooks()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			apperr.NewAppErr(apperr.INTERNAL_ERROR, fmt.Sprintf("cannot get books: %v", err)))
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetBook returns book by given ID
func (bc *bookController) GetBook(c *gin.Context) {

	var bookName, status string

	ID := uuid.MustParse(c.Param("id"))
	response, err := bc.bookService.GetBook(ID)
	if err != nil {
		bookName = "error"
		status = strconv.Itoa(http.StatusInternalServerError)
		prometheus.BookStatus.WithLabelValues(bookName, status).Inc()
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			apperr.NewAppErr(apperr.INTERNAL_ERROR, fmt.Sprintf("cannot get book: %v", err)))
		return
	}

	bookName = response.Name
	status = strconv.Itoa(http.StatusOK)
	prometheus.BookStatus.WithLabelValues(bookName, status).Inc()
	c.JSON(http.StatusOK, response)
}
