package borrows

import (
	"context"
	"fmt"
	"github.com/BieLuk/library-backend/src/db/mongo"
	"github.com/BieLuk/library-backend/src/model"
	"github.com/BieLuk/library-backend/src/utils"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	mongodb "go.mongodb.org/mongo-driver/mongo"
	"time"
)

type borrowsMongoRepository struct {
	ctx               context.Context
	borrowsCollection *mongodb.Collection
}

func NewBorrowMongoRepository(ctx context.Context) *borrowsMongoRepository {
	return &borrowsMongoRepository{
		ctx:               ctx,
		borrowsCollection: mongo.GetDB().Collection("borrows"),
	}
}

func (r *borrowsMongoRepository) CreateBorrow(borrow *model.Borrow) (*model.Borrow, error) {
	borrow.ID = utils.Pointer(uuid.New())
	borrow.CreatedAt, borrow.UpdatedAt = time.Now(), time.Now()
	_, err := r.borrowsCollection.InsertOne(r.ctx, borrow)
	if err != nil {
		return nil, fmt.Errorf("error creating book in database: %w", err)
	}

	return borrow, nil
}

func (r *borrowsMongoRepository) GetBorrowByBookID(bookID uuid.UUID) (*model.Borrow, error) {
	borrow := model.Borrow{}
	result := r.borrowsCollection.FindOne(r.ctx, bson.M{"book_id": bookID})
	if result.Err() != nil {
		return nil, fmt.Errorf("error retrieving borrow from database: %w", result.Err())
	}
	if err := result.Decode(&borrow); err != nil {
		return nil, fmt.Errorf("error decoding result from database: %w", err)
	}

	return &borrow, nil
}
