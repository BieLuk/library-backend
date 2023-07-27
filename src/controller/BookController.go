package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type BookController interface {
	GetBooks(c *gin.Context)
}

type bookController struct {
}

func NewBookController() *bookController {
	return &bookController{}
}

// GetBooks returns list of all books in library
func (bc *bookController) GetBooks(c *gin.Context) {
	books := make([]GetBookResponse, 0)
	books = append(books, GetBookResponse{
		ID:          uuid.New(),
		Name:        "Harry Potter and the Sorcerer's Stone",
		Author:      "J. K. Rowling",
		ISBN:        "9788380082113",
		Description: "An orphaned boy enrolls in a school of wizardry, where he learns the truth about himself, his family and the terrible evil that haunts the magical world.",
	})
	books = append(books, GetBookResponse{
		ID:          uuid.New(),
		Name:        "Natthuset",
		Author:      "Jo Nesbo",
		ISBN:        "9788327164001",
		Description: "In the wake of his parents’ tragic deaths in a house fire, fourteen-year-old Richard Elauved has been sent to live with his aunt and uncle in the remote, insular town of Ballantyne. Richard quickly earns a reputation as an outcast, and when a classmate named Tom goes missing, everyone suspects the new, angry boy is responsible for his disappearance.",
	})

	c.JSON(http.StatusOK, &GetBooksResponse{Books: books})
}
