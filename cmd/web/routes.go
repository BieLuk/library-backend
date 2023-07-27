package main

func (s *libraryServer) routes() {
	booksGroup := s.server.Group("/books")
	booksGroup.GET("/", s.bookController.GetBooks)
}
