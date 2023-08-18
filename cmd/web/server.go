package main

import (
	"context"
	"fmt"
	"github.com/BieLuk/library-backend/src/config"
	"github.com/BieLuk/library-backend/src/controller/book"
	"github.com/BieLuk/library-backend/src/controller/borrow"
	"github.com/BieLuk/library-backend/src/db/mongo"
	"github.com/BieLuk/library-backend/src/db/postgres"
	booksRepo "github.com/BieLuk/library-backend/src/repository/books"
	borrowsRepo "github.com/BieLuk/library-backend/src/repository/borrows"
	"github.com/BieLuk/library-backend/src/service/books"
	"github.com/BieLuk/library-backend/src/service/borrows"
	"github.com/gin-gonic/gin"
)

type libraryServer struct {
	server *gin.Engine

	bookController   book.BookController
	borrowController borrow.BorrowController
}

func runLibraryServer() {
	appConfig, err := config.LoadConfig(".")
	if err != nil {
		panic(fmt.Errorf("error occurred reading config file"))
	}

	var bookRepository booksRepo.BooksRepository
	var borrowRepository borrowsRepo.BorrowsRepository
	if config.DatabaseEngine(appConfig.DatabaseEngine) == config.DatabaseEngineMongo {
		if err := mongo.Init(context.Background(), appConfig.MongoDBURI, appConfig.MongoDBName); err != nil {
			panic(fmt.Errorf("error initializing database: %w", err))
		}
		bookRepository = booksRepo.NewBookMongoRepository(context.Background())
		borrowRepository = borrowsRepo.NewBorrowMongoRepository(context.Background())
	} else {
		if err := postgres.Init(appConfig); err != nil {
			panic(fmt.Errorf("error initializing postgres database: %w", err))
		}
		bookRepository = booksRepo.NewBookPostgresRepository()
		borrowRepository = borrowsRepo.NewBorrowPostgresRepository()
	}

	bookService := books.NewBookService(bookRepository)
	borrowService := borrows.NewBorrowService(borrowRepository)

	libServer := &libraryServer{
		server:           gin.Default(),
		bookController:   book.NewBookController(bookService),
		borrowController: borrow.NewBorrowController(borrowService),
	}

	libServer.routes()

	if err := libServer.server.Run(fmt.Sprintf(":%s", appConfig.ServerPort)); err != nil {
		panic(fmt.Errorf("error occurred running library server: %w", err))
	}
}
