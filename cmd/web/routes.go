package main

func (s *libraryServer) routes() {
	booksGroup := s.server.Group("/books")
	booksGroup.POST("/", s.bookController.CreateBook)
	booksGroup.GET("/", s.bookController.GetBooks)
	booksGroup.GET("/:id", s.bookController.GetBook)

	borrowsGroup := s.server.Group("/borrows")
	borrowsGroup.POST("/", s.borrowController.CreateBorrow)
	borrowsGroup.GET("/checkBorrow/:id", s.borrowController.IsBookBorrowed)
}
