package main

func (s *libraryServer) routes() {
	booksGroup := s.server.Group("/books")
	booksGroup.POST("/", s.bookController.CreateBook)
}
