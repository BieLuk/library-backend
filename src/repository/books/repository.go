package books

import (
	"fmt"
	"github.com/BieLuk/library-backend/src/db"
	"github.com/BieLuk/library-backend/src/model"
	"github.com/google/uuid"
)

type BooksRepository interface {
	CreateBook(book *model.Book) (*model.Book, error)
	GetBooks() ([]*model.Book, error)
	GetBook(bookID uuid.UUID) (*model.Book, error)
}

type booksRepository struct {
}

func NewBookRepository() *booksRepository {
	return &booksRepository{}
}

func (r *booksRepository) CreateBook(book *model.Book) (*model.Book, error) {
	if result := db.GetDB().Create(book); result.Error != nil {
		return nil, fmt.Errorf("error creating book in database: %w", result.Error)
	}
	return book, nil
}

func (r *booksRepository) GetBooks() ([]*model.Book, error) {
	var books []*model.Book
	if result := db.GetDB().Find(&books); result.Error != nil {
		return nil, fmt.Errorf("error retrieving books from database: %w", result.Error)
	}
	return books, nil
}

func (r *booksRepository) GetBook(bookID uuid.UUID) (*model.Book, error) {
	var book *model.Book
	if result := db.GetDB().Where("id = ?", bookID).Take(&book); result.Error != nil {
		return nil, fmt.Errorf("error retrieving book from database: %w", result.Error)
	}
	return book, nil
}
