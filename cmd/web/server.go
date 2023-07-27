package main

import (
	"fmt"
	"github.com/BieLuk/library-backend/src/config"
	"github.com/BieLuk/library-backend/src/controller/book"
	"github.com/BieLuk/library-backend/src/db"
	booksRepo "github.com/BieLuk/library-backend/src/repository/books"
	"github.com/BieLuk/library-backend/src/service/books"
	"github.com/gin-gonic/gin"
)

type libraryServer struct {
	server *gin.Engine

	bookController book.BookController
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
	bookService := books.NewBookService(bookRepository)

	libServer := &libraryServer{
		server:         gin.Default(),
		bookController: book.NewBookController(bookService),
	}

	libServer.routes()

	if err := libServer.server.Run(fmt.Sprintf(":%s", appConfig.ServerPort)); err != nil {
		panic(fmt.Errorf("error occurred running library server: %w", err))
	}
}
