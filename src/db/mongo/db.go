package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var DbName string
var client *mongo.Client

func Init(ctx context.Context, mongoURI string, dbName string) error {
	DbName = dbName
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("%s%s", mongoURI, dbName)))
	if err != nil {
		return fmt.Errorf("cannot connect mongo db: %w", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return fmt.Errorf("cannot ping mongo db: %w", err)
	}

	return nil
}

func GetClient() *mongo.Client {
	return client
}
