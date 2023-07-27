package main

import (
	"fmt"
	"github.com/BieLuk/library-backend/src/config"
	"github.com/BieLuk/library-backend/src/controller/book"
	"github.com/BieLuk/library-backend/src/controller/borrow"
	"github.com/BieLuk/library-backend/src/db"
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
	if err := db.Init(appConfig); err != nil {
		panic(fmt.Errorf("error initializing database: %w", err))
	}

	bookRepository := booksRepo.NewBookRepository()
	borrowRepository := borrowsRepo.NewBorrowRepository()
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
