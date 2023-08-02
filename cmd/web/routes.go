package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (s *libraryServer) routes() {
	booksGroup := s.server.Group("/books")
	booksGroup.POST("/", s.bookController.CreateBook)
	booksGroup.GET("/", s.bookController.GetBooks)
	booksGroup.GET("/:id", s.bookController.GetBook)

	borrowsGroup := s.server.Group("/borrows")
	borrowsGroup.POST("/", s.borrowController.CreateBorrow)
	borrowsGroup.GET("/checkBorrow/:id", s.borrowController.IsBookBorrowed)

	s.server.GET("/metrics", prometheusHandler())
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
