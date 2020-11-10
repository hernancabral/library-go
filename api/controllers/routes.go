package controllers

func (s *Server) initializeRoutes() {

	v1 := s.Router.Group("/api/v1")
	{
		// Books routes
		v1.POST("/books", s.PostBook)
		v1.GET("/books", s.GetBooks)
		v1.GET("/books/:id", s.GetBookById)
		v1.PUT("/books/:id", s.UpdateBook)
		v1.DELETE("/books/:id", s.DeleteBook)

		// Genres routes
		v1.GET("/genres", s.GetGenres)

		// Authors routes
		v1.GET("/authors", s.GetAuthors)

		// Search routes
		v1.GET("/search/publisher/:publisher", s.GetBooksByPublisher)
		v1.GET("/search/years/:years", s.GetBooksByYear)
		v1.GET("/search/publisher/:publisher/years/:years", s.GetBooksByPublisherAndYear)
		v1.GET("/search/keyword/:keyword", s.GetBooksByKeyword)
		v1.GET("/search/title/:title", s.GetBooksByTitle)
		v1.GET("/search/author/:author", s.GetBooksByAuthor)

		// Seed DB
		v1.POST("/seed", s.Seed)
	}
}
