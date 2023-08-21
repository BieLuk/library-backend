package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *libraryServer) routes() {
	booksGroup := s.server.Group("/books")
	booksGroup.POST("/", s.bookController.CreateBook)
	booksGroup.GET("/", s.bookController.GetBooks)
	booksGroup.GET("/:id", s.bookController.GetBook)
	booksGroup.PUT("/:id", s.bookController.UpdateBook)
	booksGroup.DELETE("/:id", s.bookController.DeleteBook)

	borrowsGroup := s.server.Group("/borrows")
	borrowsGroup.POST("/", s.borrowController.CreateBorrow)
	borrowsGroup.GET("/check/:id", s.borrowController.IsBookBorrowed)
	borrowsGroup.PUT("/return/", s.borrowController.ReturnBorrowedBook)

	s.server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	s.server.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
