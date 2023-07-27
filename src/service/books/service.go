package books

import (
	"fmt"
	"github.com/BieLuk/library-backend/src/dto"
	"github.com/BieLuk/library-backend/src/model"
	"github.com/BieLuk/library-backend/src/repository/books"
	"github.com/google/uuid"
)

type BookService interface {
	CreateBook(request dto.CreateBookRequest) (*dto.CreateBookResponse, error)
	GetBooks() (*dto.GetBooksResponse, error)
	GetBook(ID uuid.UUID) (*dto.GetBookResponse, error)
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

	bookEntity, err := s.bookRepository.CreateBook(bookEntity)
	if err != nil {
		return nil, fmt.Errorf("cannot create bookEntity: %w", err)
	}

	return &dto.CreateBookResponse{ID: *bookEntity.ID}, nil
}

func (s *bookService) GetBooks() (*dto.GetBooksResponse, error) {
	books, err := s.bookRepository.GetBooks()
	if err != nil {
		return nil, fmt.Errorf("cannot get books: %w", err)
	}

	response := dto.GetBooksResponse{}
	for _, book := range books {
		response.Books = append(response.Books, dto.GetBookResponse{
			ID:          book.ID,
			Name:        book.Name,
			Author:      book.Author,
			ISBN:        book.ISBN,
			Description: book.Description,
		})
	}

	return &response, nil
}

func (s *bookService) GetBook(bookID uuid.UUID) (*dto.GetBookResponse, error) {
	book, err := s.bookRepository.GetBook(bookID)
	if err != nil {
		return nil, fmt.Errorf("cannot get books: %w", err)
	}

	return &dto.GetBookResponse{
		ID:          book.ID,
		Name:        book.Name,
		Author:      book.Author,
		ISBN:        book.ISBN,
		Description: book.Description,
	}, nil
}
