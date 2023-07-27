package books

import (
	"fmt"
	"github.com/BieLuk/library-backend/src/db"
	"github.com/BieLuk/library-backend/src/model"
)

type BooksRepository interface {
	CreateBook(book *model.Book) error
}

type booksRepository struct {
}

func NewBookRepository() *booksRepository {
	return &booksRepository{}
}

func (r *booksRepository) CreateBook(book *model.Book) error {
	if result := db.GetDB().Create(book); result.Error != nil {
		return fmt.Errorf("error creating book in database: %w", result.Error)
	}
	return nil
}
