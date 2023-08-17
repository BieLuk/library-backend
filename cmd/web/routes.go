package main

import (
	"context"
	"fmt"
	"github.com/BieLuk/library-backend/src/db/mongo"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *libraryServer) routes() {
	booksGroup := s.server.Group("/books")
	booksGroup.POST("/", s.bookController.CreateBook)
	booksGroup.GET("/", s.bookController.GetBooks)
	booksGroup.GET("/:id", s.bookController.GetBook)

	borrowsGroup := s.server.Group("/borrows")
	borrowsGroup.POST("/", s.borrowController.CreateBorrow)
	borrowsGroup.GET("/checkBorrow/:id", s.borrowController.IsBookBorrowed)

	s.server.GET("/metrics", gin.WrapH(promhttp.Handler()))

	s.server.GET("/test", test)
}

func test(c *gin.Context) {
	user := bson.D{{"fullName", "User 1"}, {"age", 30}}
	usersCollection := mongo.GetClient().Database(mongo.DbName).Collection("users")
	result, err := usersCollection.InsertOne(context.Background(), user)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}
