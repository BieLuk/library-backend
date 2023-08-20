package books

import (
	"fmt"
	"github.com/BieLuk/library-backend/src/dto"
	"github.com/BieLuk/library-backend/src/model"
	"github.com/BieLuk/library-backend/src/repository/books"
	"github.com/BieLuk/library-backend/src/repository/borrows"
	"github.com/google/uuid"
)

type BookService interface {
	CreateBook(request dto.CreateBookRequest) (*dto.CreateBookResponse, error)
	GetBooks() (*dto.GetBooksResponse, error)
	GetBook(ID uuid.UUID) (*dto.GetBookResponse, error)
	UpdateBook(ID uuid.UUID, request dto.UpdateBookRequest) error
	DeleteBook(ID uuid.UUID) error
}

type bookService struct {
	bookRepository   books.BooksRepository
	borrowRepository borrows.BorrowsRepository
}

func NewBookService(bookRepository books.BooksRepository, borrowRepository borrows.BorrowsRepository) *bookService {
	return &bookService{
		bookRepository:   bookRepository,
		borrowRepository: borrowRepository,
	}
}

func (s *bookService) CreateBook(request dto.CreateBookRequest) (*dto.CreateBookResponse, error) {
	bookEntity := &model.Book{
		Name:        request.Name,
		Author:      request.Author,
		ISBN:        request.ISBN,
		Description: request.Description,
		Status:      model.BookStatusActive,
	}

	bookEntity, err := s.bookRepository.CreateBook(bookEntity)
	if err != nil {
		return nil, fmt.Errorf("cannot create book: %w", err)
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
			ID:          *book.ID,
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
		ID:          *book.ID,
		Name:        book.Name,
		Author:      book.Author,
		ISBN:        book.ISBN,
		Description: book.Description,
	}, nil
}

func (s *bookService) UpdateBook(bookID uuid.UUID, request dto.UpdateBookRequest) error {
	bookEntity := &model.Book{
		DBEntity: model.DBEntity{
			ID: &bookID,
		},
		Name:        request.Name,
		Author:      request.Author,
		ISBN:        request.ISBN,
		Description: request.Description,
	}

	err := s.bookRepository.UpdateBook(bookEntity)
	if err != nil {
		return fmt.Errorf("cannot update book: %w", err)
	}

	return nil
}

func (s *bookService) DeleteBook(bookID uuid.UUID) error {
	borrowsNotBrought, err := s.borrowRepository.GetBorrowsNotBroughtByBookID(bookID)
	if err != nil {
		return fmt.Errorf("cannot get not brought borrows: %w", err)
	}
	if len(borrowsNotBrought) > 0 {
		return fmt.Errorf("cannot delete book, there are not brought borrows")
	}
	if err := s.bookRepository.DeleteBook(bookID); err != nil {
		return fmt.Errorf("cannot delete book: %w", err)
	}

	return nil
}
