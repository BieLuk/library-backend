package books

import (
	"fmt"
	"github.com/BieLuk/library-backend/src/db/postgres"
	"github.com/BieLuk/library-backend/src/model"
	"github.com/google/uuid"
)

type booksPostgresRepository struct {
}

func NewBookPostgresRepository() *booksPostgresRepository {
	return &booksPostgresRepository{}
}

func (r *booksPostgresRepository) CreateBook(book *model.Book) (*model.Book, error) {
	if result := postgres.GetDB().Create(book); result.Error != nil {
		return nil, fmt.Errorf("error creating book in database: %w", result.Error)
	}
	return book, nil
}

func (r *booksPostgresRepository) GetBooks() ([]*model.Book, error) {
	var books []*model.Book
	if result := postgres.GetDB().Where("status=?", model.BookStatusActive).Find(&books); result.Error != nil {
		return nil, fmt.Errorf("error retrieving books from database: %w", result.Error)
	}
	return books, nil
}

func (r *booksPostgresRepository) GetBook(bookID uuid.UUID) (*model.Book, error) {
	var book *model.Book
	if result := postgres.GetDB().Where("id = ?", bookID).Where("status=?", model.BookStatusActive).Take(&book); result.Error != nil {
		return nil, fmt.Errorf("error retrieving book from database: %w", result.Error)
	}
	return book, nil
}

func (r *booksPostgresRepository) UpdateBook(book *model.Book) error {
	if result := postgres.GetDB().Updates(book); result.Error != nil {
		return fmt.Errorf("error updating book: %w", result.Error)
	}

	return nil
}

func (r *booksPostgresRepository) DeleteBook(bookID uuid.UUID) error {
	if result := postgres.GetDB().Model(model.Book{}).Where("id=?", bookID).
		UpdateColumn("status", model.BookStatusDeleted); result.Error != nil {
		return fmt.Errorf("error deleting book: %w", result.Error)
	}

	return nil
}
