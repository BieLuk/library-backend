package main

func (s *libraryServer) routes() {
	booksGroup := s.server.Group("/books")
	booksGroup.POST("/", s.bookController.CreateBook)
	booksGroup.GET("/", s.bookController.GetBooks)
	booksGroup.GET("/:id", s.bookController.GetBook)
}
