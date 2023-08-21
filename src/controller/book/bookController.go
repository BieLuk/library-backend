package book

import (
	"fmt"
	"github.com/BieLuk/library-backend/src/apperr"
	"github.com/BieLuk/library-backend/src/dto"
	"github.com/BieLuk/library-backend/src/service/books"
	"github.com/BieLuk/library-backend/src/validate"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type BookController interface {
	CreateBook(c *gin.Context)
	GetBooks(c *gin.Context)
	GetBook(c *gin.Context)
	UpdateBook(c *gin.Context)
	DeleteBook(c *gin.Context)
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
// @Summary      Create book
// @Description  Create book
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param        createBookRequest body dto.CreateBookRequest true "createBookRequest"
// @Success      200  {object}  dto.CreateBookResponse
// @Failure      400  {object}  apperr.AppErr
// @Router       /books/ [post]
func (bc *bookController) CreateBook(c *gin.Context) {
	var request dto.CreateBookRequest
	if err := validate.BindAndValidateJson(c, &request); err != nil {
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
// @Summary      Get all books
// @Description  Get list of all books in library
// @Tags         Books
// @Accept       json
// @Produce      json
// @Success      200  {object}  dto.GetBooksResponse
// @Failure      400  {object}  apperr.AppErr
// @Router       /books/ [get]
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
// @Summary      Get book by id
// @Description  Get single book by given id
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param        id   path string  true  "Book ID"
// @Success      200  {object}  dto.GetBookResponse
// @Failure      400  {object}  apperr.AppErr
// @Router       /books/{id} [get]
func (bc *bookController) GetBook(c *gin.Context) {
	ID := uuid.MustParse(c.Param("id"))
	response, err := bc.bookService.GetBook(ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			apperr.NewAppErr(apperr.INTERNAL_ERROR, fmt.Sprintf("cannot get book: %v", err)))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateBook updates book data
// @Summary      Update book
// @Description  Update book data
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param        id   path string  true  "Book ID"
// @Param        updateBookRequest body dto.UpdateBookRequest true "updateBookRequest"
// @Success      200
// @Failure      400  {object}  apperr.AppErr
// @Router       /books/{id} [put]
func (bc *bookController) UpdateBook(c *gin.Context) {
	ID := uuid.MustParse(c.Param("id"))
	var request dto.UpdateBookRequest
	if err := validate.BindAndValidateJson(c, &request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			apperr.NewAppErr(apperr.BAD_REQUEST, fmt.Sprintf("cannot unmarshall request object: %v", err)))
		return
	}
	err := bc.bookService.UpdateBook(ID, request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			apperr.NewAppErr(apperr.INTERNAL_ERROR, fmt.Sprintf("cannot update book: %v", err)))
		return
	}

	c.Status(http.StatusOK)
}

// DeleteBook deletes book
// @Summary      Delete book
// @Description  Delete book
// @Tags         Books
// @Accept       json
// @Produce      json
// @Param        id   path string  true  "Book ID"
// @Success      200
// @Failure      400  {object}  apperr.AppErr
// @Router       /books/{id} [delete]
func (bc *bookController) DeleteBook(c *gin.Context) {
	ID := uuid.MustParse(c.Param("id"))

	err := bc.bookService.DeleteBook(ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			apperr.NewAppErr(apperr.INTERNAL_ERROR, fmt.Sprintf("cannot delete book: %v", err)))
		return
	}

	c.Status(http.StatusOK)
}
