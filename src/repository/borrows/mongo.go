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

func (r *borrowsMongoRepository) GetBorrowsByBookID(bookID uuid.UUID) ([]*model.Borrow, error) {
	var borrows []*model.Borrow
	filter := bson.M{"book_id": bookID}
	result, err := r.borrowsCollection.Find(r.ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error retrieving borrow from database: %w", err)
	}

	if err := result.All(r.ctx, &borrows); err != nil {
		return nil, fmt.Errorf("error marshalling results from DB: %w", err)
	}

	return borrows, nil
}

func (r *borrowsMongoRepository) GetBorrowsNotBroughtByBookID(bookID uuid.UUID) ([]*model.Borrow, error) {
	var borrows []*model.Borrow
	filter := bson.M{"book_id": bookID, "brought_date": nil}
	result, err := r.borrowsCollection.Find(r.ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error retrieving borrow from database: %w", err)
	}

	if err := result.All(r.ctx, &borrows); err != nil {
		return nil, fmt.Errorf("error marshalling results from DB: %w", err)
	}

	return borrows, nil
}

func (r *borrowsMongoRepository) UpdateBorrowBroughtDateByBookID(bookID uuid.UUID, broughtDate time.Time) error {
	_, err := r.borrowsCollection.UpdateOne(r.ctx, bson.M{"book_id": bookID}, bson.M{
		"$set": model.Borrow{
			BroughtDate: &broughtDate,
			BookID:      bookID,
			DBEntity: model.DBEntity{
				UpdatedAt: time.Now(),
			}},
	})
	if err != nil {
		return fmt.Errorf("error updating borrow: %w", err)
	}
	return nil

}
