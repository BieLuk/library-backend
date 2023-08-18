package books

import (
	"context"
	"fmt"
	mongodb "github.com/BieLuk/library-backend/src/db/mongo"
	"github.com/BieLuk/library-backend/src/model"
	"github.com/BieLuk/library-backend/src/utils"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type booksMongoRepository struct {
	ctx             context.Context
	booksCollection *mongo.Collection
}

func NewBookMongoRepository(ctx context.Context) *booksMongoRepository {
	return &booksMongoRepository{
		ctx:             ctx,
		booksCollection: mongodb.GetDB().Collection("books"),
	}
}

func (r *booksMongoRepository) CreateBook(book *model.Book) (*model.Book, error) {
	book.ID = utils.Pointer(uuid.New())
	book.CreatedAt, book.UpdatedAt = time.Now(), time.Now()
	_, err := r.booksCollection.InsertOne(r.ctx, book)
	if err != nil {
		return nil, fmt.Errorf("error creating book in database: %w", err)
	}

	return book, nil
}

func (r *booksMongoRepository) GetBooks() ([]*model.Book, error) {
	result, err := r.booksCollection.Find(r.ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error retrieving books from database: %w", err)
	}

	var books []*model.Book
	if err := result.All(r.ctx, &books); err != nil {
		return nil, fmt.Errorf("error marshalling results from DB: %w", err)
	}

	return books, nil
}

func (r *booksMongoRepository) GetBook(bookID uuid.UUID) (*model.Book, error) {
	book := model.Book{}
	result := r.booksCollection.FindOne(r.ctx, bson.M{"_id": bookID})
	if result.Err() != nil {
		return nil, fmt.Errorf("error retrieving book from database: %w", result.Err())
	}
	if err := result.Decode(&book); err != nil {
		return nil, fmt.Errorf("error decoding result from database: %w", err)
	}

	return &book, nil
}

func (r *booksMongoRepository) UpdateBook(book *model.Book) error {
	book.UpdatedAt = time.Now()
	_, err := r.booksCollection.UpdateOne(r.ctx, bson.M{"_id": book.ID}, bson.M{
		"$set": book,
	})
	if err != nil {
		return fmt.Errorf("error updating book: %w", err)
	}
	return nil
}

func (r *booksMongoRepository) DeleteBook(bookID uuid.UUID) error {
	_, err := r.booksCollection.UpdateOne(r.ctx, bson.M{"_id": bookID}, bson.M{
		"$set": model.Book{
			Status: model.BookStatusDeleted,
			DBEntity: model.DBEntity{
				ID:        &bookID,
				UpdatedAt: time.Now(),
			}},
	})
	if err != nil {
		return fmt.Errorf("error updating book: %w", err)
	}
	return nil
}
