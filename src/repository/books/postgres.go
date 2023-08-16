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
	if result := postgres.GetDB().Find(&books); result.Error != nil {
		return nil, fmt.Errorf("error retrieving books from database: %w", result.Error)
	}
	return books, nil
}

func (r *booksPostgresRepository) GetBook(bookID uuid.UUID) (*model.Book, error) {
	var book *model.Book
	if result := postgres.GetDB().Where("id = ?", bookID).Take(&book); result.Error != nil {
		return nil, fmt.Errorf("error retrieving book from database: %w", result.Error)
	}
	return book, nil
}
