package books

import (
	"fmt"
	"github.com/BieLuk/library-backend/src/dto"
	"github.com/BieLuk/library-backend/src/model"
	"github.com/BieLuk/library-backend/src/repository/books"
)

type BookService interface {
	CreateBook(request dto.CreateBookRequest) (*dto.CreateBookResponse, error)
}

type bookService struct {
	bookRepository books.BooksRepository
}

func NewBookService(bookRepository books.BooksRepository) *bookService {
	return &bookService{
		bookRepository: bookRepository,
	}
}

func (s *bookService) CreateBook(request dto.CreateBookRequest) (*dto.CreateBookResponse, error) {
	bookEntity := &model.Book{
		Name:        request.Name,
		Author:      request.Author,
		ISBN:        request.ISBN,
		Description: request.Description,
	}

	if err := s.bookRepository.CreateBook(bookEntity); err != nil {
		return nil, fmt.Errorf("cannot create bookEntity: %w", err)
	}

	return &dto.CreateBookResponse{ID: *bookEntity.ID}, nil
}
