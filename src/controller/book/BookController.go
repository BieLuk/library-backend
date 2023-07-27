package book

import (
	"fmt"
	"github.com/BieLuk/library-backend/src/apperr"
	"github.com/BieLuk/library-backend/src/dto"
	"github.com/BieLuk/library-backend/src/service/books"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookController interface {
	CreateBook(c *gin.Context)
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
	}

	response, err := bc.bookService.CreateBook(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			apperr.NewAppErr(apperr.BAD_REQUEST, fmt.Sprintf("cannot unmarshall request object: %v", err)))

	}

	c.JSON(http.StatusOK, response)
}

//
//// GetBooks returns list of all books in library
//func (bc *bookController) GetBooks(c *gin.Context) {
//	books := make([]GetBookResponse, 0)
//	books = append(books, GetBookResponse{
//		ID:          uuid.New(),
//		Name:        "Harry Potter and the Sorcerer's Stone",
//		Author:      "J. K. Rowling",
//		ISBN:        "9788380082113",
//		Description: "An orphaned boy enrolls in a school of wizardry, where he learns the truth about himself, his family and the terrible evil that haunts the magical world.",
//	})
//	books = append(books, GetBookResponse{
//		ID:          uuid.New(),
//		Name:        "Natthuset",
//		Author:      "Jo Nesbo",
//		ISBN:        "9788327164001",
//		Description: "In the wake of his parents’ tragic deaths in a house fire, fourteen-year-old Richard Elauved has been sent to live with his aunt and uncle in the remote, insular town of Ballantyne. Richard quickly earns a reputation as an outcast, and when a classmate named Tom goes missing, everyone suspects the new, angry boy is responsible for his disappearance.",
//	})
//
//	c.JSON(http.StatusOK, &GetBooksResponse{Books: books})
//}
