package books

import (
	"github.com/BieLuk/library-backend/src/model"
	"github.com/google/uuid"
)

type BooksRepository interface {
	CreateBook(book *model.Book) (*model.Book, error)
	GetBooks() ([]*model.Book, error)
	GetBook(bookID uuid.UUID) (*model.Book, error)
	UpdateBook(book *model.Book) error
	DeleteBook(bookID uuid.UUID) error
}
